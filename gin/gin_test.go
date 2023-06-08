package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"sokrUrl/database"
	"testing"
)

func TestPostUrl(t *testing.T) {
	database.NewPostgreSQLClient()
	var url = "/urls"
	r := gin.Default()
	r.POST(url, postUrl)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/url", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)

	var fullUrl = "https://habr.com/ru/articles/441842/"
	req2, _ := http.NewRequest(http.MethodPost, "/urls?url="+fullUrl, nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req2)
	//var sokrUrl = w.Body.String()
	assert.Equal(t, http.StatusCreated, w.Code)
	//	assert.Equal(t, "{\"Сокращенный url\":\"https://habr.com/lZtxtwlYmL\"}", sokrUrl)
}

func TestGetUrl(t *testing.T) {
	database.NewPostgreSQLClient()
	var url = "/urls"
	r := gin.Default()
	r.GET(url, getFullUrl)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/url", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)

	var shortUrl = "https://habr.com/THFNxlW7QG"
	req2, _ := http.NewRequest(http.MethodGet, "/urls?url="+shortUrl, nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req2)
	var fullUrl = w.Body.String()
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"Полный url\":\"https://habr.com/ru/articles/441842/\"}", fullUrl)
}
