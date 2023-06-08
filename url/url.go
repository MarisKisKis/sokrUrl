package url

import (
	"log"
	"math/rand"
	"net/url"
)

type Url struct {
	FullUrl  string
	ShortUrl string
}

func CreateShortUrl(fullUrl string) string {
	var shortUrl string
	u, err := url.Parse(fullUrl)
	if err != nil {
		log.Fatal(err)
	}
	var scheme = u.Scheme
	var host = u.Host
	var seq = randSeq(10)
	shortUrl = scheme + "://" + host + "/" + seq
	return shortUrl
}

var alphabet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(b)
}
