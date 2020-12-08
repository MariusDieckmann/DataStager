package cmd

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/MariusDieckmann/DataStager/loader.go"
	"github.com/spf13/cobra"
)

// uploadCommand represents the uploadCommand command
var uploadCommand = &cobra.Command{
	Use:   "upload",
	Short: "Uploads a given set of files to a given bucket with a given basekey",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		bucket, err := cmd.Flags().GetString("bucket")
		if err != nil {
			log.Fatalln(err.Error())
		}

		baseKey, err := cmd.Flags().GetString("basekey")
		if err != nil {
			log.Fatalln(err.Error())
		}

		filePaths, err := cmd.Flags().GetStringSlice("files")
		if err != nil {
			log.Fatalln(err.Error())
		}

		loader := loader.InitLoader(CustomS3Endpoint)
		err = loader.UploadFiles(bucket, baseKey, filePaths)
		if err != nil {
			log.Fatalln(err.Error())
		}

	},
}

func init() {
	rootCmd.AddCommand(uploadCommand)
	uploadCommand.Flags().StringP("bucket", "b", "", "Bucket to upload the data to")
	uploadCommand.Flags().StringP("basekey", "k", "", "Basekey to upload the data to")
	uploadCommand.Flags().StringSliceP("files", "f", []string{}, "List of files to upload")

	err := uploadCommand.MarkFlagRequired("bucket")
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = uploadCommand.MarkFlagRequired("basekey")
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = uploadCommand.MarkFlagRequired("files")
	if err != nil {
		log.Fatalln(err.Error())
	}

}

func countLines(filePath string) int {
	var buffer bytes.Buffer

	cmd := exec.Command("wc", "-l", filePath)
	cmd.Stderr = os.Stderr
	cmd.Stdout = &buffer
	err := cmd.Run()
	if err != nil {
		log.Fatalln(err.Error())
	}

	stdoutResponse := buffer.String()
	splittedResponse := strings.Split(stdoutResponse, " ")

	contigCount, err := strconv.Atoi(splittedResponse[0])
	if err != nil {
		log.Fatalln(err.Error())
	}

	return contigCount
}
