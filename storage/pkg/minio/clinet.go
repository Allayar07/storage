package minio

import (
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewMinioClient(endpoint, accesskeyId, secretaccesskeyId string) (*minio.Client, error) {
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accesskeyId, secretaccesskeyId, ""),
		Secure: false,
		Region: "us-east-1",
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create minio client. err: %w", err)
	}

	return minioClient, nil
}
