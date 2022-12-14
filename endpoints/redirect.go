package endpoints

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"urlShortener/dal"

	"github.com/gin-gonic/gin"
)

type Redirect struct {
	DbClient *dal.DynamoDBClient
}

func NewRedirect() Redirect {
	db := dal.GetDbClient()

	return Redirect{DbClient: db}
}

func (r *Redirect) Handler(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid id input, expected int")
		return
	}

	redirectUrl, err := r.DbClient.GetRedirect(idInt)
	if err != nil {
		// TODO: implement not found response
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, "Error reading from db")
		return
	}

	*redirectUrl = addProtocolIfNotExists(*redirectUrl)

	c.Redirect(http.StatusMovedPermanently, *redirectUrl)
}

func addProtocolIfNotExists(redirectUrl string) string {
	if strings.HasPrefix(redirectUrl, "http") {
		return redirectUrl
	}

	return "https://" + redirectUrl
}
