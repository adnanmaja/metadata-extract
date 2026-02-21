package api

import (
	"bytes"
	"io"
	"net/http"

	metadata "github.com/adnanmaja/metadata-extract/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Upload(c *gin.Context) {
	var req struct {
		BlobURL      string `json:"blobUrl"`
		OriginalName string `json:"originalName"`
	}

	if err := c.BindJSON(&req); err != nil || req.BlobURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing blobUrl"})
		return
	}

	resp, err := http.Get(req.BlobURL)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to fetch blob"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(500, gin.H{"error": "blob fetch failed"})
		return
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to read blob"})
		return
	}

	reader := bytes.NewReader(data)

	exif, metadata, config, err := metadata.FileHandler(reader)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"basic": gin.H{
			"FileName":         req.OriginalName,
			"FileSize":         float64(len(data)) / 1048576.0,
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

func GetUploadURL(c *gin.Context) {
	blobName := uuid.New().String() + ".jpg"

	uploadURL, blobURL, err := metadata.GenerateUploadURL(blobName)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"uploadUrl": uploadURL,
		"blobUrl":   blobURL,
	})
}
