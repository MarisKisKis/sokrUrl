package gin

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	url2 "net/url"
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
	a := assert.New(t)
	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}
	var actual Data
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}
	assert.Equal(t, http.StatusCreated, w.Code)
	u, err := url2.Parse(actual.Url)
	if err != nil {
		log.Fatal(err)
	}
	var scheme = u.Scheme
	var host = u.Host
	assert.Equal(t, "https", scheme)
	assert.Equal(t, "habr.com", host)
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
	a := assert.New(t)
	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}
	var actual Data
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "https://habr.com/ru/articles/441842/", actual.Url)
}

type Data struct {
	Url string `json:"url"`
}
