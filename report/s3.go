package report

import (
	"bytes"
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const (
	AWS_S3_REGION         = "ap-south-1"
	AWS_S3_BUCKET         = "AWS_S3_BUCKET"
	AWS_ACCESS_KEY_ID     = "AWS_ACCESS_KEY_ID"
	AWS_SECRET_ACCESS_KEY = "AWS_SECRET_ACCESS_KEY/7HWKVUGdhdY"
)

var filePath string = "/Users/vijeth.ag/go/src/confluence/report.csv"

func UploadToS3() error {

	s3Config := &aws.Config{
		Region:      aws.String(AWS_S3_REGION),
		Credentials: credentials.NewStaticCredentials(AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, ""),
	}
	s3Session := session.New(s3Config)

	uploader := s3manager.NewUploader(s3Session)

	upFile, err := os.Open(filePath)
	if err != nil {
		log.Println("err opening file", err)
		return err
	}
	defer upFile.Close()

	upFileInfo, _ := upFile.Stat()
	var fileSize int64 = upFileInfo.Size()
	fileBuffer := make([]byte, fileSize)
	upFile.Read(fileBuffer)

	input := &s3manager.UploadInput{
		Bucket:      aws.String(AWS_S3_BUCKET),   // bucket's name
		Key:         aws.String("report.csv"),    // files destination location
		Body:        bytes.NewReader(fileBuffer), // content of the file
		ContentType: aws.String("file/csv"),      // content type
	}
	output, err := uploader.UploadWithContext(context.Background(), input)

	if err != nil {
		log.Println("err uploading", err)
		return err
	}
	log.Println("output", output)
	return nil

}
