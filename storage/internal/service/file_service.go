package service

import (
	"context"
	"crypto/sha1"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
	"io"
	"os"
	"storage/internal/repository"
)

const (
	from     = "deoj559@gmail.com"
	password = "srsafyahrpgmtlgs"
	smtpHost = "smtp.gmail.com"
	smtpPort = 587
)

type FileService struct {
	repo   repository.FileStorage
	client *minio.Client
}

func NewFileService(repo repository.FileStorage, client *minio.Client) *FileService {
	return &FileService{
		repo:   repo,
		client: client,
	}
}

func (s *FileService) Upload(ctx context.Context, bucketName, fileName, ContentType, email string, fileSize int64, reader io.Reader) (string, error) {
	exist, errBucketExist := s.client.BucketExists(ctx, bucketName)
	if errBucketExist != nil || !exist {
		err := s.client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{
			Region:        "us-east-1",
			ObjectLocking: true,
		})
		if err != nil {
			return "", err
		}
	}

	info, err := s.client.PutObject(ctx, bucketName, fileName, reader, fileSize, minio.PutObjectOptions{
		ContentType: ContentType,
	})

	if err != nil {
		return "", fmt.Errorf("failed to upload file. Err: %w", err)
	}
	Key := GenerateKey(fileName)
	err = s.repo.UploadDB(fileName, Key)
	if err != nil {
		return "", err
	}
	msg := fmt.Sprintf("http://localhost:8080/api/get/%s", Key)
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", email)
	//m.SetHeader("Subject", "from somebody")
	m.SetBody("text", msg)

	d := gomail.NewDialer(smtpHost, smtpPort, from, password)

	if err := d.DialAndSend(m); err != nil {
		return "", fmt.Errorf("failed to sending file link. Err: %w", err)
	}
	logrus.Printf("successfully uploaded %s of size %d\n", fileName, info.Size)

	return msg, nil
}

func (s *FileService) Download(ctx context.Context, key string) (*minio.Object, string, error) {
	var BucketName string = "test"
	FileName, err := s.repo.GetKey(key)

	if err != nil {
		return nil, "", err
	}

	object, err := s.client.GetObject(ctx, BucketName, FileName, minio.GetObjectOptions{})
	if err != nil {
		return nil, "", err
	}
	localfile, err := os.Create("downloadedFiles/" + FileName)
	if err != nil {
		return nil, "", err
	}
	defer localfile.Close()

	stat, err := object.Stat()
	if err != nil {
		return nil, "", err
	}
	if _, err := io.CopyN(localfile, object, stat.Size); err != nil {
		return nil, "", err
	}

	return object, FileName, nil
}

func (s *FileService) Delete(ctx context.Context, bucketName, key string) error {
	fileName, err := s.repo.GetKey(key)
	if err != nil {
		return err
	}
	err = s.client.RemoveObject(ctx, bucketName, fileName, minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}
	if err := s.repo.DeleteFile(key); err != nil {
		return err
	}
	return nil
}

func GenerateKey(key string) string {
	hash := sha1.New()
	hash.Write([]byte(key))

	return fmt.Sprintf("%x", hash.Sum([]byte("some")))
}
