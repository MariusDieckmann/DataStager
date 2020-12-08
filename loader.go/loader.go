package loader

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

//Loader Data loader to stage data from and to object storage
type Loader struct {
	Downloader *s3manager.Downloader
	Uploader   *s3manager.Uploader
}

//InitLoader Initiates a new data loader
func InitLoader(customS3Endpoint string) Loader {

	config := &aws.Config{}

	if customS3Endpoint != "" {
		config.Endpoint = aws.String(customS3Endpoint)
		config.Region = aws.String("RegionOne")
	}

	session := session.Must(session.NewSession(config))
	downloader := s3manager.NewDownloader(session)
	uploader := s3manager.NewUploader(session)

	loader := Loader{
		Downloader: downloader,
		Uploader:   uploader,
	}

	return loader
}
