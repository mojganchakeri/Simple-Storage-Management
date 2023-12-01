package api

import (
	"fmt"
	"os"

	"store_service/api/controller"
	"store_service/api/middleware"
	"store_service/configs"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

func SetupServer() {
	gin.SetMode(gin.ReleaseMode)
	os.Setenv("GIN_MODE", "release")
	r := gin.New()

	// Swagger handler
	if configs.SwaggerEnable {
		docs.SwaggerInfo.BasePath = ""
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}

	// API handler
	r.Use(gin.Recovery())
	// r.GET("/ping", handlers.PingHandlerGET)
	authorized := r.Group("/api")
	authorized.Use(middleware.CheckUser)
	{
		v1 := authorized.Group("/v1")
		{
			v1.POST("/upload", controller.UploadFile)
			v1.POST("/retrieve", controller.RetrieveFile)
		}
	}
	logrus.Info(fmt.Sprintf("service starts at %s", configs.Env.ServerAddress))
	r.Run(configs.Env.ServerAddress)
}
