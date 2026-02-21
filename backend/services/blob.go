package metadata

import (
	"fmt"
	"os"
	"time"

	azblob "github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/sas"
)

func GenerateUploadURL(blobName string) (uploadURL string, blobURL string, err error) {
	accountName := os.Getenv("STORAGE_ACCOUNT_BAME")
	accountKey := os.Getenv("STORAGE_ACCOUNT_KEY")

	cred, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		return "", "", err
	}

	perms := sas.BlobPermissions{
		Create: true,
		Write:  true,
	}

	startTime := time.Now().UTC().Add(-5 * time.Minute)
	expiryTime := startTime.Add(20 * time.Minute)

	sasQueryParams, err := sas.BlobSignatureValues{
		Protocol:      sas.ProtocolHTTPS,
		StartTime:     startTime,
		ExpiryTime:    expiryTime,
		ContainerName: "upload",
		BlobName:      blobName,
		Permissions:   perms.String(),
	}.SignWithSharedKey(cred)
	if err != nil {
		return "", "", err
	}

	uploadURL = fmt.Sprintf(
		"https://%s.blob.core.windows.net/upload/%s?%s",
		accountName,
		blobName,
		sasQueryParams.Encode(),
	)

	blobURL = fmt.Sprintf(
		"https://%s.blob.core.windows.net/upload/%s",
		accountName,
		blobName,
	)

	return uploadURL, blobURL, nil
}
