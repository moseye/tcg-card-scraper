package helpers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"tcg-scraper/config"
	"tcg-scraper/server"
)

func NewServer(db *gorm.DB) *server.Server {
	s := &server.Server{
		Echo:   echo.New(),
		DB:     db,
		Config: config.NewConfig(),
	}

	return s
}
