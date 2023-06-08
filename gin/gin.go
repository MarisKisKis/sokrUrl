package gin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sokrUrl/database"
	"sokrUrl/url"
)

func home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Get url",
	})
	return
}

func postUrl(c *gin.Context) {
	fullUrl := c.Query("url")
	var shortUrl = url.CreateShortUrl(fullUrl)
	var url database.Url
	url.ShortUrl = shortUrl
	url.FullUrl = fullUrl
	res2, err2 := database.CreateUrl(&url)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err2,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"Сокращенный url": res2,
	})
	return
}

func getFullUrl(c *gin.Context) {
	var shortUrl = c.Query("url")
	url, err := database.FullUrl(shortUrl)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "url not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Полный url": url,
	})
	return
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", home)
	r.GET("/urls", getFullUrl)
	r.POST("/urls", postUrl)
	return r
}
