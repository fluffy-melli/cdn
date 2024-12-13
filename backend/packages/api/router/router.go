package router

import (
	_ "cdn-module/docs"
	"cdn-module/packages/api/ui"
	"cdn-module/packages/cache"
	"cdn-module/packages/logger"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Setup(memory *cache.Safe) *gin.Engine {
	router := gin.New()
	router.Use(logger.Logger())
	router.Use(gin.Recovery())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/content/*filename", ui.SendFile)
	router.GET("/content-list", ui.FileList)
	router.POST("/file/upload", ui.UploadFile)

	router.GET("/upload", func(c *gin.Context) {
		c.File("./asset/upload.html")
	})

	router.Static("/static", "./asset")

	router.LoadHTMLGlob("templates/*")
	router.GET("/view/*filename", Embed)

	return router
}
