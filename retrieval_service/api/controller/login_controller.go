package controller

import (
	"net/http"
	"retreival_service/configs"
	"retreival_service/internal"
	"retreival_service/internal/models"
	"retreival_service/repository"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @Summary Login
// @Description Login to get access and refresh token
// @Tags User
// @Accept json
// @Produce json
// @Router /api/v1/login [post]
// @Param body body models.LoginRequest true "request body"
func LoginController(ctx *gin.Context) {
	var request models.LoginRequest

	err := ctx.ShouldBind(&request)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	// get user by username
	user, err := repository.DBClient.GetUserByUsername(request.Username)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusNotFound, models.ErrorResponse{Error: "User not found"})
		return
	}

	// compare password hashes
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		logrus.Error("password is not correct")
		ctx.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "Invalid credentials"})
		return
	}

	// create access token
	accessToken, err := internal.CreateAccessToken(user, configs.Env.AccessTokenSecret, configs.Env.AccessTokenExpiryHour)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	// create refresh token
	refreshToken, err := internal.CreateRefreshToken(user, configs.Env.RefreshTokenSecret, configs.Env.RefreshTokenExpiryHour)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	// add accessToken to whiteList
	internal.AddAccessTokenToWhiteList(accessToken)

	loginResponse := models.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	ctx.JSON(http.StatusOK, loginResponse)
}
