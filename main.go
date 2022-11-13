package main

import (
	"net/http"
	"urlShortener/endpoints"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	short := endpoints.NewShort()
	redirect := endpoints.NewRedirect()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Service is working",
		})
	})
	r.POST("/short", short.Handler)
	r.GET("/:id", redirect.Handler)
	r.Run()
}
