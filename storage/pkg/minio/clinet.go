package minio

import (
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewMinioClient(endpoint, accessKeyId, secretKeyId string) (*minio.Client, error) {
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyId, secretKeyId, ""),
		Secure: false,
		Region: "us-east-1",
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create minio client. err: %w", err)
	}

	return minioClient, nil
}
