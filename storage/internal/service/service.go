package service

import (
	"context"
	"github.com/minio/minio-go/v7"
	"io"
	"storage/internal/model"
	"storage/internal/repository"
)

type Authorization interface {
	Create(u model.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type FileStorage interface {
	Upload(ctx context.Context, bucketName, fileName, ContentType, email string, fileSize int64, reader io.Reader) (string, error)
	Download(ctx context.Context, key string) (*minio.Object, string, error)
	Delete(ctx context.Context, bucketName, key string) error
}

type Service struct {
	Authorization
	FileStorage
}

func NewService(repo *repository.Repository, client *minio.Client) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		FileStorage:   NewFileService(repo.FileStorage, client),
	}
}
