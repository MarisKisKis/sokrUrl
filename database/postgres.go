package database

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var db *gorm.DB
var err error

type Url struct {
	ID       uint   `gorm:"primary_key"`
	FullUrl  string `gorm:"unique"`
	ShortUrl string `gorm:"unique"`
}

func getEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	return os.Getenv(key)
}

func NewPostgreSQLClient() {
	var (
		host     = getEnvVariable("DB_HOST")
		port     = getEnvVariable("DB_PORT")
		user     = getEnvVariable("DB_USER")
		dbname   = getEnvVariable("DB_NAME")
		password = getEnvVariable("DB_PASSWORD")
	)

	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		user,
		dbname,
		password,
	)

	db, err = gorm.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(Url{})
}

func CreateUrl(a *Url) (string, error) {
	res := db.Create(a)
	if res.RowsAffected == 0 {
		return a.ShortUrl, errors.New("url not created")
	}
	return a.ShortUrl, nil
}

func FullUrl(shortUrl string) (*string, error) {
	var url Url
	res := db.Where(&Url{ShortUrl: shortUrl}).Find(&url)
	if res.RowsAffected == 0 {
		return nil, errors.New("url not found")
	}
	return &url.FullUrl, nil
}
