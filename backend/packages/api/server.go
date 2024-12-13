package api

import (
	"cdn-module/config"
	_ "cdn-module/docs"
	"cdn-module/packages/api/router"
	"cdn-module/packages/api/ui"
	"cdn-module/packages/cache"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run(cf *config.CDN_CONFIG_JSON, memory *cache.Safe) {
	gin.SetMode(gin.ReleaseMode)
	router := router.Setup(memory)
	config := cors.Config{
		AllowOrigins:     cf.Server.Cors.Config.AllowOrigins,
		AllowMethods:     cf.Server.Cors.Config.AllowMethods,
		AllowHeaders:     cf.Server.Cors.Config.AllowHeaders,
		ExposeHeaders:    cf.Server.Cors.Config.ExposeHeaders,
		AllowCredentials: cf.Server.Cors.Config.AllowCredentials,
	}
	ui.INIT(cf, memory)
	router.MaxMultipartMemory = cf.Server.MaxFSMB * 1024 * 1024
	router.Use(cors.New(config))
	router.Run(fmt.Sprintf(":%d", cf.Server.Port))
}
