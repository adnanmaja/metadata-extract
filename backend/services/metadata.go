package metadata

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"github.com/evanoberholster/imagemeta"
	"github.com/evanoberholster/imagemeta/exif2"
)

func ExtractMetadata(f io.ReadSeeker) (exif2.Exif, error) {
	e, err := imagemeta.Decode(f)
	if err != nil {
		return exif2.Exif{}, fmt.Errorf("imagemeta decode failed: %w", err)
	}
	return e, nil
}

func FileHandler(fileHeader multipart.FileHeader) (exif2.Exif, error) {
	src, err := fileHeader.Open()
	if err != nil {
		return exif2.Exif{}, fmt.Errorf("Fileheader failed: %w", err)
	}
	defer src.Close()

	tempFile, err := os.CreateTemp("", "upload-*.tmp")
	if err != nil {
		return exif2.Exif{}, fmt.Errorf("Fileheader failed: %w", err)
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	io.Copy(tempFile, src)
	tempFile.Seek(0, io.SeekStart)

	exif, _ := ExtractMetadata(tempFile)
	return exif, nil
}
