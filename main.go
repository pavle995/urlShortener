package main

import (
	"net/http"
	"urlShortener/endpoints"

	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	r := gin.Default()

	initLog()

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
	r.Run()
}

func initLog() {
	file, err := os.OpenFile("/var/log/urlShortner/urlShortner.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("Failed to open log file" + err.Error())
	}
	log.SetOutput(file)
}
