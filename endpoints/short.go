package endpoints

import (
	"net/http"
	"time"
	"urlShortener/dal"
	"urlShortener/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Short struct {
	DbClient *dal.DynamoDBClient
}

func (s *Short) Handler(c *gin.Context) {
	fullUrl := c.PostForm("url")
	uuid := uuid.New()

	// hardcoded for now
	ttl := 86400
	created := time.Now()

	url := models.Url{
		Id:        uuid.ID(),
		Url:       fullUrl,
		CreatedAt: created,
		TTL:       ttl,
	}
	err := s.DbClient.InsertNewRecord(&url)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"shortUrl": url.Id,
			"message":  "inserted",
		})
	}
}
