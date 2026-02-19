package api

import (
	"net/http"

	metadata "github.com/adnanmaja/metadata-extract/services"
	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "something"})
		return
	}
	exif, _ := metadata.FileHandler(*fileHeader)

	c.JSON(200, gin.H{
		"basic": gin.H{
			"FileName":         fileHeader.Filename,
			"FileSize":         fileHeader.Size,
			"FileType":         exif.ImageType,
			"DateTimeOriginal": exif.DateTimeOriginal(),
			"ImageWidth":       exif.ImageWidth,
			"ImageHeight":      exif.ImageHeight,
		},
		"camera": gin.H{
			"Make":         exif.CameraMake,
			"Model":        exif.CameraModel,
			"LensInfo":     exif.LensModel,
			"FNumber":      exif.FNumber,
			"ExposureTime": exif.ExposureTime,
			"ISO":          exif.ISOSpeed,
			"FocalLength":  exif.FocalLength,
		},
		"location": gin.H{
			"GPSLatitude":  exif.GPS.Latitude(),
			"GPSLongitude": nil,
			"GPSAltitude":  exif.GPS.Altitude(),
			"City":         nil,
			"Country":      nil,
		},
	})
}
