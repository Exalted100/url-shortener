package routes

import (
	"fmt"
	"net/http"
	"time"
	"url-shortener/src/config"
	"url-shortener/src/db"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/teris-io/shortid"
)

type urlShortenRequest struct {
	URL string `json:"url"`
	TTL int64  `json:"ttl"`
}

type urlShortenResponse struct {
	URL          string `json:"url"`
	ShortenedUrl string `json:"shortenedUrl"`
}

func getShortURLHandler(c *gin.Context) {
	var currentShortenRequest urlShortenRequest

	if err := c.BindJSON(&currentShortenRequest); err != nil {
		return
	}

	url := currentShortenRequest.URL
	ttl := currentShortenRequest.TTL

	uniqueID, err := createUniqueId()

	if err != nil && err != redis.Nil {
		errString := "Could not generate a unique ID"
		fmt.Println("Error: Could not generate a unique ID for url shortener")
		c.JSON(http.StatusInternalServerError, errString)
	}

	err = db.Client.Set(uniqueID, url, time.Duration(ttl*int64(time.Second))).Err()

	if err != nil {
		errString := "Could not save unique ID to redis database"
		fmt.Println("Error: Could not save unique ID to redis database for url shortener")
		c.JSON(http.StatusInternalServerError, errString)
	}

	c.JSON(200, gin.H{
		"code": "SUCCESS",
		"data": urlShortenResponse{
			URL:          url,
			ShortenedUrl: config.ConfigVariables.AppHost + "/" + uniqueID,
		},
	})
}

func redirectIfURLFound(c *gin.Context) {

	urlShortCode := c.Param("url")
	actualUrl, err := db.Client.Get(urlShortCode).Result()
	if actualUrl == "" || err != nil {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	} else {
		c.Redirect(http.StatusFound, actualUrl)
		c.Abort()
	}
}

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    "SUCCESS",
		"message": "The RESTful server is OK!",
	})
}

func createUniqueId() (string, error) {
	uniqueID, uniqueIdErr := shortid.Generate()
	if uniqueIdErr != nil {
		fmt.Println("Could not generate unique ID with the package")
	}
	_, err := db.Client.Get(uniqueID).Result()
	if err != redis.Nil {
		return createUniqueId()
	}
	return uniqueID, err
}
