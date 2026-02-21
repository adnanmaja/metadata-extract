package metadata

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"os"
	"time"

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

func FileHandler(src io.ReadSeeker) (
	exif2.Exif,
	map[string]interface{},
	image.Config,
	error,
) {
	totalStart := time.Now()

	copyStart := time.Now()
	tempFile, err := os.CreateTemp("", "upload-*.tmp")
	if err != nil {
		return exif2.Exif{}, nil, image.Config{}, err
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	if _, err := io.Copy(tempFile, src); err != nil {
		return exif2.Exif{}, nil, image.Config{}, err
	}
	fmt.Fprintf(gin.DefaultWriter, "[TIMER] File copy: %v\n", time.Since(copyStart))

	exifStart := time.Now()
	tempFile.Seek(0, io.SeekStart)
	exif, _ := extractMetadata(tempFile)
	fmt.Fprintf(gin.DefaultWriter, "[TIMER] Imagemeta exif: %v\n", time.Since(exifStart))

	goExifStart := time.Now()
	tempFile.Seek(0, io.SeekStart)
	additionalMetadata, _ := extractMetadataGoExif(tempFile)
	fmt.Fprintf(gin.DefaultWriter, "[TIMER] Go Exif: %v\n", time.Since(goExifStart))

	osStart := time.Now()
	tempFile.Seek(0, io.SeekStart)
	config, _ := extractMetadataOs(tempFile)
	fmt.Fprintf(gin.DefaultWriter, "[TIMER] Image.Config (OS): %v\n", time.Since(osStart))

	fmt.Fprintf(gin.DefaultWriter, "[TIMER] Total FileHandler: %v\n", time.Since(totalStart))
	return exif, additionalMetadata, config, nil
}

func extractMetadataGoExif(f io.ReadSeeker) (map[string]interface{}, error) {
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
