package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/evanoberholster/imagemeta"
	"github.com/evanoberholster/imagemeta/exif2"
	"github.com/gin-gonic/gin"
)

func main() {
	post()
}

func extract(f io.ReadSeeker) (exif2.Exif, error) {
	e, err := imagemeta.Decode(f)
	if err != nil {
		return exif2.Exif{}, fmt.Errorf("imagemeta decode failed: %w", err)
	}
	return e, nil
}

func post() {
	router := gin.Default()
	router.POST("/upload", func(c *gin.Context) {
		fileHeader, err := c.FormFile("image")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "something"})
			return
		}
		src, err := fileHeader.Open()
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to open file"})
			return
		}
		defer src.Close()

		tempFile, err := os.CreateTemp("", "upload-*.tmp")
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to create temp file"})
			return
		}
		defer os.Remove(tempFile.Name())
		defer tempFile.Close()

		io.Copy(tempFile, src)
		tempFile.Seek(0, io.SeekStart)

		exif, _ := extract(tempFile)
		c.JSON(200, gin.H{
			"Author: ":      exif.Artist,
			"Camera make: ": exif.CameraMake,
			"F stop: ":      exif.FNumber,
		})
	})
	router.Run(":8080")
}
