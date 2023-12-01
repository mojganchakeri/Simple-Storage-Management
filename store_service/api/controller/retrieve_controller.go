package controller

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"store_service/configs"
	"store_service/internal"
	"store_service/internal/models"
	"store_service/repository"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @Summary Retreive file
// @Description Retreive file base search
// @Tags Retreive
// @Accept json
// @Produce json
// @Router /api/v1/retreive [post]
// @Param  body body models.RequestRetreive true "request_body"
func RetrieveFile(ctx *gin.Context) {
	var request models.RequestRetreive

	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	filesPath, ferr := repository.DBClient.SearchFileInDB(request.Name, request.Tag)
	if ferr != nil {
		logrus.Error("search query to find file is failed")
		ctx.JSON(http.StatusBadGateway, models.ErrorResponse{Error: ferr.Error()})
		return

	}

	zipPath, zerr := readFiles(filesPath)
	if zerr != nil {
		logrus.Error("reading files is failed")
		ctx.JSON(http.StatusBadGateway, models.ErrorResponse{Error: zerr.Error()})
		return
	}

	zipFile, er := os.Open(zipPath)
	if er != nil {
		// return 404 HTTP response code for File not found
		logrus.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Error: er.Error()})
		return
	}
	fileHeader := make([]byte, 512)
	zipFile.Read(fileHeader)
	fileType := http.DetectContentType(fileHeader) // set the type
	fileInfo, _ := zipFile.Stat()
	fileSize := fileInfo.Size()

	//Transmit the headers
	ctx.Writer.Header().Set("Expires", "0")
	ctx.Writer.Header().Set("Content-type", "application/zip")
	ctx.Writer.Header().Set("Content-Control", "private, no-transform, no-store, must-revalidate")
	ctx.Writer.Header().Set("Content-Disposition", "attachment; filename="+filepath.Base(zipPath))
	ctx.Writer.Header().Set("Content-Type", fileType)
	ctx.Writer.Header().Set("Filename", filepath.Base(zipPath))

	ctx.Writer.Header().Set("Content-Length", strconv.FormatInt(fileSize, 10))
	http.ServeFile(ctx.Writer, ctx.Request, zipPath)
	zipFile.Seek(0, 0)

	os.Remove(zipPath)

}

func readFiles(filesPath []string) (string, error) {
	fixFiles := []string{}

	for _, fpath := range filesPath {
		fileBytes, ferr := os.ReadFile(fpath)
		if ferr != nil {
			logrus.Error(fmt.Sprintf("file %s is intrupted", fpath))
			continue
		}

		// Encrypt the file contents
		decreptedData := internal.Decrypt(string(fileBytes))

		// Get main path
		// fpathSplits := strings.Split(fpath, "/")
		// mainPath := configs.Env.StoragePath
		// if len(fpathSplits) >0 {
		// 	fileId := fpathSplits[len(fpathSplits)-1]
		// 	mainPath = fpath[:len(fileId)-1]
		// 	println("main path len: ", mainPath)
		// }
		tempFile, err := os.Create(fpath)

		if err != nil {
			logrus.Error(err.Error())
		}

		defer tempFile.Close()
		_, err = tempFile.Write(decreptedData)
		if err != nil {
			logrus.Error(err.Error())
			fixFiles = append(fixFiles, fpath)
		}
	}

	zipPath, zerr := makeZip(fixFiles)
	if zerr != nil {
		return zipPath, zerr
	}
	return zipPath, nil

}

func makeZip(filesPath []string) (string, error) {
	//Create a new zip archive and named archive.zip
	zipPath := fmt.Sprintf("%sexport.zip", configs.Env.StoragePath)
	archive, err := os.Create(zipPath)
	if err != nil {
		logrus.Error(err)
		return zipPath, err
	}
	defer archive.Close()

	//Create a new zip writer
	zipWriter := zip.NewWriter(archive)

	// Add files to the zip archive
	fixedFilePath := []string{}

	for _, fname := range filesPath {
		f1, err := os.Open(fname)
		if err != nil {
			logrus.Error(fmt.Sprintf("Reading file %s is distrupted", fname))
			continue
		}
		defer f1.Close()
		w1, err := zipWriter.Create(fname)
		if err != nil {
			logrus.Error(fmt.Sprintf("Reading file %s is distrupted", fname))
			continue
		}
		if _, err := io.Copy(w1, f1); err != nil {
			logrus.Error(err)
			continue
		}

		fixedFilePath = append(fixedFilePath, fname)

	}

	deleteFile(fixedFilePath)

	zipWriter.Close()

	return zipPath, nil
}

func deleteFile(filesPath []string) error {
	for _, fpath := range filesPath {
		rerr := os.Remove(fpath)
		if rerr != nil {
			logrus.Error(fmt.Sprintf("file %s is not deleted", fpath))
		}
	}

	return repository.DBClient.DeleteFile(filesPath)
}
