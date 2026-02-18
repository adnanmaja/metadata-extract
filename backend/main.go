package main

import (
	api "github.com/adnanmaja/metadata-extract/api"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/upload", api.Upload)
	router.Run(":8080")
}
