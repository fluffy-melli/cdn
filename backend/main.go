package main

import (
	"cdn-module/config"
	"cdn-module/packages/api"
	"cdn-module/packages/cache"
	"cdn-module/packages/logger"
	"fmt"
)

func main() {
	fmt.Println(logger.Fg_BrightWhite + logger.LineString(" [LOAD - Cache] ", 70) + logger.Reset)
	df := config.NEW_CDN_CONFIG_JSON()
	ch := cache.ALLCache(df)
	fmt.Println(logger.Fg_BrightWhite + logger.LineString("", 70) + logger.Reset)
	api.Run(df, ch)
}
