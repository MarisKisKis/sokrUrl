package main

import (
	"sokrUrl/database"
	"sokrUrl/gin"
)

func init() {
	database.NewPostgreSQLClient()
}

func main() {
	r := gin.SetupRouter()
	r.Run("localhost:8080")
}
