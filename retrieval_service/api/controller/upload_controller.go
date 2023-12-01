package controller

import (
	"fmt"
	"retreival_service/configs"
	"retreival_service/internal"

	"github.com/gin-gonic/gin"
)

func UploadController(ctx *gin.Context) {
	storeServiceAddress := fmt.Sprintf("%s:%s", configs.Env.StoreServiceHost, configs.Env.StoreServicePort)
	internal.ReverseProxy(ctx, storeServiceAddress, "/api/v1/upload")
}
