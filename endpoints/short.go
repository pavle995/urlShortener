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

type Body struct {
	Url string `form:"url"`
}

func NewShort() Short {
	db := dal.GetDbClient()

	return Short{DbClient: db}
}

func (s *Short) Handler(c *gin.Context) {
	var fullUrl string
	var body Body
	if c.ShouldBind(&body) == nil {
		fullUrl = body.Url
	}
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
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"shortUrl": url.Id,
		"message":  "inserted",
	})
}
