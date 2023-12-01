package controller

import (
	"fmt"
	"net/http"
	"retreival_service/internal/models"
	"retreival_service/repository"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func RegisterController(ctx *gin.Context) {
	var request models.LoginRequest

	err := ctx.ShouldBind(&request)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	// check user duplicate in db
	dbUser, err := repository.DBClient.GetUserByUsername(request.Username)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	if dbUser.ID != "" {
		err = fmt.Errorf("user %s is existed", request.Username)
		logrus.Error(err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	// save user in database
	var user models.User
	user.ID = uuid.New().String()
	user.Username = request.Username

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 14)
	user.Password = string(passwordHash)

	if err := repository.DBClient.AddNewUser(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "internal server error"})
		return
	}

	ctx.JSON(http.StatusCreated, models.Response{Message: "user is created"})
}
