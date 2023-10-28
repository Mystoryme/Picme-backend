package minio

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"

	mod "picme-backend/modules"
)

func Init() {
	endpoint := "minio.bsthun.com"
	accessKeyID := "ExffndqFKmbQpUhe"
	secretAccessKey := "JXrDtva30KDLJpMspeg81qRYM2QiI6Jl"
	useSSL := true

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		log.Fatalln(err)
	}

	mod.Minio = minioClient
}
