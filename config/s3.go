package config

import (
	"os"

	"github.com/gofiber/storage/minio"
)

func S3() *minio.Storage {
	endpoint := os.Getenv("S3_ADDRESS")
	accessKeyID := os.Getenv("S3_USER")
	secretAccessKey := os.Getenv("S3_PASSWORD")
	useSSL := false

	// Initialize MinIO client
	MinioClient := minio.New(minio.Config{
		Endpoint: endpoint,
		Secure:   useSSL,
		Bucket:   "uploads",
		Credentials: minio.Credentials{
			AccessKeyID:     accessKeyID,
			SecretAccessKey: secretAccessKey,
		},
	})
	return MinioClient
}
