package controller

import (
	"net/http"
	"retreival_service/internal"
	"retreival_service/internal/models"

	"github.com/gin-gonic/gin"
)

func LogoutController(ctx *gin.Context) {

	// remove old accessToken from whiteList
	internal.RemoveAccessTokenFromWhiteList(ctx.Request.Header.Get("Authorization"))

	ctx.JSON(http.StatusOK, models.Response{Message: "logout successful"})
}
