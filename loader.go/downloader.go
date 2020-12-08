package loader

import (
	"errors"
	"log"
	"os"
	"path"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

//ErrInequalKeyAndBucketLen Error to indicate inequal number of buckets and keys
var ErrInequalKeyAndBucketLen = errors.New("Foo")

//DownloadFiles Downloads the files into a directory
func (loader *Loader) DownloadFiles(buckets []string, keys []string, targetDir string) error {
	if len(buckets) != len(keys) {
		return ErrInequalKeyAndBucketLen
	}

	for i := range buckets {
		currentBucket := buckets[i]
		currentKey := keys[i]

		err := loader.downloadS3(currentBucket, currentKey, targetDir)
		if err != nil {
			log.Println(err.Error())
			return err
		}

	}

	return nil
}

func (loader *Loader) downloadS3(bucket string, key string, targetDir string) error {
	_, filename := path.Split(key)
	file, err := os.Create(path.Join(targetDir, filename))

	defer file.Close()

	if err != nil {
		log.Println(err.Error())
		return err
	}

	objectInput := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	_, err = loader.Downloader.Download(file, objectInput)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
