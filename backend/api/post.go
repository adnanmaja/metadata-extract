package api

import (
	"net/http"

	metadata "github.com/adnanmaja/metadata-extract/services"
	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	exif, metadata, config, _ := metadata.FileHandler(*fileHeader)

	c.JSON(200, gin.H{
		"basic": gin.H{
			"FileName":         fileHeader.Filename,
			"FileSize":         float64(fileHeader.Size) / 1048576.0,
			"FileType":         exif.ImageType,
			"DateTimeOriginal": exif.DateTimeOriginal(),
			"ImageWidth":       config.Width,
			"ImageHeight":      config.Height,
			"Software":         metadata["Software"],
		},
		"camera": gin.H{
			"Make":         metadata["Make"],
			"Model":        metadata["Model"],
			"LensInfo":     exif.LensModel,
			"FNumber":      exif.FNumber,
			"ExposureTime": exif.ExposureTime,
			"ISO":          exif.ISOSpeed,
			"FocalLength":  exif.FocalLength,
			"CameraSerial": metadata["BodySerialNumber"],
			"LensSerial":   metadata["LensSerialNumber"],
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
