package controller

import (
	"net/http"
	"retreival_service/configs"
	"retreival_service/internal"
	"retreival_service/internal/models"
	"retreival_service/repository"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @Summary Refresh
// @Description Refresh to get access and refresh token
// @Tags User
// @Accept json
// @Produce json
// @Router /api/v1/refresh [post]
// @Param body body models.RefreshRequest true "request body"
func RefreshController(ctx *gin.Context) {
	var request models.RefreshRequest

	err := ctx.ShouldBind(&request)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	id, err := internal.ExtractIDFromToken(request.RefreshToken, configs.Env.RefreshTokenSecret)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "User not fount"})
		return
	}

	// get user by id
	user, err := repository.DBClient.GetUserByID(id)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusNotFound, models.ErrorResponse{Error: "User not found"})
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

	// remove old accessToken from whiteList
	internal.RemoveAccessTokenFromWhiteList(ctx.Request.Header.Get("Authorization"))

	// add new accessToken to whiteList
	internal.AddAccessTokenToWhiteList(accessToken)

	loginResponse := models.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	ctx.JSON(http.StatusOK, loginResponse)
}
