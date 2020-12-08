package loader

import (
	"log"
	"os"
	"path"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

//UploadFiles Uploads the files to the given object storage location
func (loader *Loader) UploadFiles(bucket string, baseKey string, filepaths []string) error {
	for _, filepath := range filepaths {
		err := loader.uploadFile(bucket, baseKey, filepath)
		if err != nil {
			log.Println(err.Error())
			return err
		}
	}

	return nil
}

func (loader *Loader) uploadFile(bucket string, baseKey string, filepath string) error {
	_, filename := path.Split(filepath)
	uploadKey := path.Join(baseKey, filename)

	file, err := os.Open(filepath)
	defer file.Close()

	if err != nil {
		log.Println(err.Error())
		return err
	}

	uploadInput := &s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(uploadKey),
		Body:   file,
	}

	_, err = loader.Uploader.Upload(uploadInput)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
