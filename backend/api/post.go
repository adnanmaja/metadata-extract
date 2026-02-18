package api

import (
	"net/http"

	metadata "github.com/adnanmaja/metadata-extract/services"
	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	fileHeader, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "something"})
		return
	}
	exif, _ := metadata.FileHandler(*fileHeader)

	c.JSON(200, gin.H{
		"Author: ":      exif.Artist,
		"Camera make: ": exif.CameraMake,
		"F stop: ":      exif.FNumber,
	})
}
