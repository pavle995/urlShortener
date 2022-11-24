package main

import (
	"net/http"
	"urlShortener/endpoints"

	"os"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	r := gin.Default()
	r.Use(CORSMiddleware())

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
	log.Fatal(autotls.Run(r, "urlshortnerapi.pavlekosutic.com"))
}

func initLog() {
	file, err := os.OpenFile("/var/log/urlShortner/urlShortner.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("Failed to open log file" + err.Error())
	}
	log.SetOutput(file)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
