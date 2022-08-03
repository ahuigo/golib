package aws

import (
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/spf13/viper"
	"github.com/ahuigo/glogger"
)

var logger = glogger.Glogger

// Client aws
type Client struct {
	sess *session.Session
}

// NewDefaultClient constructor
func NewDefaultClient() *Client {
	creds := credentials.NewStaticCredentialsFromCreds(
		credentials.Value{
			AccessKeyID:     viper.GetString("aws.accessKeyID"),
			SecretAccessKey: viper.GetString("aws.secretAccessKey"),
		})
	awsConfig := &aws.Config{
		Credentials:      creds,
		Endpoint:         aws.String(viper.GetString("aws.endpoint")),
		Region:           aws.String(viper.GetString("aws.region")),
		S3ForcePathStyle: aws.Bool(true),
	}
	return &Client{
		sess: session.Must(session.NewSession(awsConfig)),
	}
}

// Upload aws
func (c *Client) Upload(input io.Reader, bucket string, key string) error {
	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(c.sess)

	// Upload the file to S3.
	uploadResult, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: &bucket,
		Key:    &key,
		Body:   input,
	})
	logger.Debug("uploadResult", uploadResult)
	return err
}

// Download from aws
func (c *Client) Download(bucket, key string) (output []byte, err error) {
	buffer := new(aws.WriteAtBuffer)
	downloader := s3manager.NewDownloader(c.sess)
	_, err = downloader.Download(buffer,
		&s3.GetObjectInput{
			Bucket: &bucket,
			Key:    &key,
		})
	if err != nil {
		return
	}
	output = buffer.Bytes()
	return
}
