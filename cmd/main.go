package main

import (
	"fmt"
	application "tcg-scraper"
	"tcg-scraper/config"
	"tcg-scraper/docs"
)

//	@title			tcg scraper API
//	@version		0.0.1
//	@description	Go POC of backend API, will be rust in the future

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization

// @BasePath	/
func main() {
	cfg := config.NewConfig()

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.Port)

	application.Start(cfg)
}
