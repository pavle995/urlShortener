package main

import (
	"fmt"
	"net/http"
	"urlShortener/config"
	"urlShortener/endpoints"
	"urlShortener/middlewares"

	"os"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())

	cfg := config.LoadConfig()
	fmt.Println(cfg.Service.BaseUrl)

	initLog(cfg)

	short := endpoints.NewShort()
	redirect := endpoints.NewRedirect()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Service is working",
		})
	})
	r.POST("/short", short.Handler)
	r.GET("/:id", redirect.Handler)
	log.Info("Starting server...")
	log.Fatal(autotls.Run(r, cfg.Service.BaseUrl))
}

func initLog(c *config.Config) {
	file, err := os.OpenFile(c.Service.LogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("Failed to open log file" + err.Error())
	}
	log.SetOutput(file)
}
