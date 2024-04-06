package api

import (
	_ "github/NotBoring/api/docs" // swag
	v1 "github/NotBoring/api/handlers/v1"
	"github/NotBoring/config"
	"github/NotBoring/pkg/logger"
	"github/NotBoring/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Option ...
type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
}

// New ...
func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
	})
	
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"*"}
	corsConfig.AllowBrowserExtensions = true
	corsConfig.AllowMethods = []string{"*"}
	router.Use(cors.New(corsConfig))
  

	api := router.Group("/v1")
	// activity
	api.GET("/activity", handlerV1.GetActivity)	



	url := ginSwagger.URL("swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
