package controller

import (
	"fmt"
	"net/http"
	"retreival_service/configs"

	"github.com/gin-gonic/gin"
)

func UploadController(ctx *gin.Context) {
	storeServiceAddress := fmt.Sprintf("http://%s:%s/api/v1/upload", configs.Env.StoreServiceHost, configs.Env.StoreServicePort)
	ctx.Redirect(http.StatusMovedPermanently, storeServiceAddress)
}
