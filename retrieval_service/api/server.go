package api

import (
	"fmt"
	"os"

	"retreival_service/api/controller"
	"retreival_service/api/middleware"
	"retreival_service/configs"

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
	api := r.Group("/api")
	v1 := api.Group("/v1")

	v1.POST("/register", controller.RegisterController)
	v1.POST("/login", controller.LoginController)
	

	auth := v1.Group("")
	auth.Use(middleware.JwtAuthMiddleware(configs.Env.AccessTokenSecret))
	{
		auth.POST("/refresh", controller.RefreshController)
		auth.POST("/logout", controller.LogoutController)

		store := auth.Group("/store")
		store.POST("/upload", controller.UploadController)
		store.POST("/retrieve", controller.RetreiveController)
	}

	logrus.Info(fmt.Sprintf("service starts at %s", configs.Env.ServerAddress))
	r.Run(configs.Env.ServerAddress)
}
