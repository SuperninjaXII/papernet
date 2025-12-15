package controllers

import (
	"context"
	"log"
	"papernet/config"
	"time"
)

func generatePresignedURL(objectName, folder string, expiry time.Duration) (string, error) {
	bucketName := "uploads"
	minioClient := config.S3()

	// Ensure the bucket exists
	ctx := context.Background()
	objectPath := folder + "/" + objectName
	// Generate a pre-signed URL
	presignedURL, err := minioClient.Conn().PresignedGetObject(ctx, bucketName, objectPath, expiry, nil)
	if err != nil {
		log.Println("Error generating pre-signed URL:", err)
		return "", err
	}

	return presignedURL.String(), nil
}
