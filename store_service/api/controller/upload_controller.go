package controller

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"store_service/configs"
	"store_service/internal"
	"store_service/internal/models"
	"store_service/repository"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// @Summary Upload file
// @Description Upload file for user-id
// @Tags Store
// @Accept json
// @Produce json
// @Router /api/v1/upload [post]
// @Param    user-id          header               string true "user-id"
// @Param    file             formData true               true "file"
// @Param    name 			  formData true               true  "name"
// @Param    type 			  formData true               true  "type"
// @Param    tag 			  formData true               true  "tag"
func UploadFile(ctx *gin.Context) {
	fmt.Println("upload file ....................")
	var request models.RequestStore

	err := ctx.ShouldBind(&request)
	if err != nil {
		println("hereeeeeeeeeeeeeeeeeeee ", err.Error())
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	// Call controller function
	fileGormObj, allTagGormObj, fileTagGormObj, errMsg := StoreMultiPartFile(ctx)
	if errMsg != "" {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Error: errMsg})
		return
	}

	// insert to db
	derr := repository.DBClient.StoreFile(fileGormObj, allTagGormObj, fileTagGormObj)
	if derr != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Error: derr.Error()})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{Message: "Insert file in database successfully"})

}

func StoreMultiPartFile(ctx *gin.Context) (models.FileGorm, []models.TagGorm, []models.FileTagGorm, string) {

	var fileGormObj models.FileGorm
	var allTagGormObj []models.TagGorm
	var fileTagGormObj []models.FileTagGorm

	// Get file from form-data
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		logrus.Error(err.Error())
		return fileGormObj, allTagGormObj, fileTagGormObj, "Input file in form-data was intrupted"
	}

	// Generate file path
	// fileNameVal := fileHeader.Filename
	filePathVal := fmt.Sprintf("%s%s", configs.Env.StoragePath, strings.Replace(uuid.New().String(), "-", "", -1))

	// Check storage path exists or not
	filePathInfo, err := os.Stat(configs.Env.StoragePath)
	if err != nil {
		merr := os.Mkdir(configs.Env.StoragePath, 0755)
		if merr != nil {
			logrus.Error(merr.Error())
			return fileGormObj, allTagGormObj, fileTagGormObj, "Storge path not exists"
		}
	}
	if filePathInfo.Size() > int64(configs.Env.MaxStorageSize*1024*1024) {
		return fileGormObj, allTagGormObj, fileTagGormObj, "Storage path capacity exceeds max size"
	}

	// Check file size
	maxFileSize := int64(configs.Env.MaxFileSize * 1024 * 1024)
	if fileHeader.Size > maxFileSize {
		return fileGormObj, allTagGormObj, fileTagGormObj, "File size exceeds limit"
	}

	//Open received file
	csvFileToImport, err := fileHeader.Open()
	if err != nil {
		logrus.Error(err.Error())
		return fileGormObj, allTagGormObj, fileTagGormObj, "Input file in form-data was intrupted"
	}
	defer csvFileToImport.Close()

	//Create temp file
	tempFile, err := os.CreateTemp(configs.Env.StoragePath, strings.Replace(uuid.New().String(), "-", "", -1))

	if err != nil {
		logrus.Error(err.Error())
		return fileGormObj, allTagGormObj, fileTagGormObj, "Input file in form-data was intrupted"
	}

	defer tempFile.Close()

	//Write data from received file to temp file
	fileBytes, err := io.ReadAll(csvFileToImport)
	if err != nil {
		logrus.Error(err.Error())
		return fileGormObj, allTagGormObj, fileTagGormObj, "Input file in form-data was intrupted"
	}

	// Encrypt the file contents
	encryptedData, err := internal.Encrypt(fileBytes)
	if err != nil {
		logrus.Error(err.Error())
		return fileGormObj, allTagGormObj, fileTagGormObj, "Encrypting file is intrupted"
	}

	_, err = tempFile.Write(encryptedData)
	if err != nil {
		logrus.Error(err.Error())
		return fileGormObj, allTagGormObj, fileTagGormObj, "Storing file is intrupted"
	}

	reqTag := ctx.PostForm("tag")
	tagsArr := strings.Split(reqTag, ",")
	fileGormObj = models.FileGorm{
		ID:        uuid.New().String(),
		Name:      ctx.PostForm("name"),
		Type:      ctx.PostForm("type"),
		FilePath:  filePathVal,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	for _, tag := range tagsArr {
		tagId, terr := repository.DBClient.GetTagId(tag)
		if terr != nil {
			logrus.Error("search query to get tag ID is failed")
		}

		if tagId == "" {
			tagId = uuid.New().String()
			allTagGormObj = append(allTagGormObj, models.TagGorm{
				ID:    tagId,
				Value: tag,
			})
		}

		fileTagGormObj = append(fileTagGormObj, models.FileTagGorm{
			ID:     uuid.New().String(),
			TagId:  tagId,
			FileId: fileGormObj.ID,
		})
	}

	return fileGormObj, allTagGormObj, fileTagGormObj, ""
}
