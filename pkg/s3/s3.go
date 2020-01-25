package s3

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// Connection is an object used to perform S3 transactions
type Connection struct {
	s3       *s3.S3
	uploader *s3manager.Uploader
}

// New creates an instance of an S3 Connection
func New(c ConnInfo) (Connection, error) {
	cfg := aws.NewConfig().
		WithCredentials(credentials.NewStaticCredentials(c.PublicKey, c.SecretKey, c.Token)).
		WithRegion(c.Region).
		WithS3ForcePathStyle(true)

	session, err := session.NewSession(cfg)
	if err != nil {
		return Connection{}, fmt.Errorf("Failed to start AWS session with provided config [ %w ]", err)
	}

	return Connection{
		s3:       s3.New(session),
		uploader: s3manager.NewUploader(session),
	}, nil
}
