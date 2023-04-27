package main

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"net/http"
	"os"
)

type AwsRepository struct {
	config *AwsConfig
}

func NewAwsRepo(c *AwsConfig) *AwsRepository {
	return &AwsRepository{config: c}
}

func (r *AwsRepository) createSession() *session.Session {
	ses, err := session.NewSession(&aws.Config{
		Region: aws.String(r.config.Region),
		Credentials: credentials.NewStaticCredentials(
			r.config.SecretId,
			r.config.SecretKey,
			""),
	})
	if err != nil {
		panic(err)
	}
	return ses
}

func (r *AwsRepository) SaveFile(fullFilename string, fileData []byte) error {
	s := r.createSession()
	size := len(fileData)
	_, err := s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(r.config.Bucket),
		Key:                  aws.String(fullFilename),
		Body:                 bytes.NewReader(fileData),
		ContentLength:        aws.Int64(int64(size)),
		ContentType:          aws.String(http.DetectContentType(fileData)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})
	return err
}

func (r *AwsRepository) DownloadFile(object, filePathToDownload string) error {
	file, err := os.Create(filePathToDownload)
	if err != nil {
		return err
	}
	defer file.Close()

	s := r.createSession()
	numBytes, err := s3manager.NewDownloader(s).Download(file, &s3.GetObjectInput{
		Bucket: aws.String(r.config.Bucket),
		Key:    aws.String(object),
	})
	if err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("Downloaded %s with %v bytes", file.Name(), numBytes))
	return nil
}

func (r *AwsRepository) ListObjectsFromBucket() []*s3.Object {
	s := r.createSession()
	resp, err := s3.New(s).ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(r.config.Bucket)})
	if err != nil {
		panic(err)
	}
	return resp.Contents
}
