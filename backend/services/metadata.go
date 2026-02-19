package metadata

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"mime/multipart"
	"os"

	"github.com/dsoprea/go-exif/v3"
	"github.com/evanoberholster/imagemeta"
	"github.com/evanoberholster/imagemeta/exif2"
	"github.com/gin-gonic/gin"
)

func extractMetadata(f io.ReadSeeker) (exif2.Exif, error) {
	e, err := imagemeta.Decode(f)
	if err != nil {
		return exif2.Exif{}, fmt.Errorf("imagemeta decode failed: %w", err)
	}
	return e, nil
}

func FileHandler(fileHeader multipart.FileHeader) (exif2.Exif, map[string]interface{}, image.Config, error) {
	src, err := fileHeader.Open()
	if err != nil {
		return exif2.Exif{}, nil, image.Config{}, fmt.Errorf("Fileheader failed: %w", err)
	}
	defer src.Close()

	tempFile, err := os.CreateTemp("", "upload-*.tmp")
	if err != nil {
		return exif2.Exif{}, nil, image.Config{}, fmt.Errorf("Fileheader failed: %w", err)
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	io.Copy(tempFile, src)
	tempFile.Seek(0, io.SeekStart)

	exif, _ := extractMetadata(tempFile)
	tempFile.Seek(0, io.SeekStart)
	additionalMetadata, _ := extractMetadataWithGoExif(tempFile)
	tempFile.Seek(0, io.SeekStart)
	config, _ := extractMetadataOs(tempFile)
	return exif, additionalMetadata, config, nil
}

func extractMetadataWithGoExif(f io.ReadSeeker) (map[string]interface{}, error) {
	data, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	rawExif, err := exif.SearchAndExtractExif(data)
	if err != nil {
		return nil, fmt.Errorf("exif search failed: %w", err)
	}

	options := &exif.ScanOptions{}
	entries, _, err := exif.GetFlatExifData(rawExif, options)
	if err != nil {
		return nil, fmt.Errorf("failed to flatten exif data: %w", err)
	}

	results := make(map[string]interface{})
	for _, entry := range entries {
		fmt.Fprintf(gin.DefaultWriter, "Path: %s | TagName: %s | Value: %v\n", entry.IfdPath, entry.TagName, entry.Value)
		results[entry.TagName] = fmt.Sprintf("%v", entry.Value)
	}

	return results, nil
}

func extractMetadataOs(f io.ReadSeeker) (image.Config, error) {
	config, _, err := image.DecodeConfig(f)
	return config, err
}
