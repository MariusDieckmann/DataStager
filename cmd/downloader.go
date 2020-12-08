package cmd

import (
	"log"

	"github.com/MariusDieckmann/DataStager/loader.go"
	"github.com/spf13/cobra"
)

// downloadCommand represents the download command
var downloadCommand = &cobra.Command{
	Use:   "download",
	Short: "Downloads the specified data from object storage (S3) to the given targetDir",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		buckets, err := cmd.Flags().GetStringSlice("buckets")
		if err != nil {
			log.Fatalln(err.Error())
		}

		keys, err := cmd.Flags().GetStringSlice("keys")
		if err != nil {
			log.Fatalln(err.Error())
		}

		if len(buckets) != len(keys) {
			log.Fatalln("number of buckets and keys must match")
		}

		targetDir, err := cmd.Flags().GetString("targetDir")
		if err != nil {
			log.Fatalln(err.Error())
		}

		loader := loader.InitLoader(CustomS3Endpoint)
		err = loader.DownloadFiles(buckets, keys, targetDir)
		if err != nil {
			log.Fatalln(err.Error())
		}

	},
}

func init() {
	rootCmd.AddCommand(downloadCommand)
	downloadCommand.Flags().StringSliceP("buckets", "b", []string{}, "List of buckets to download data")
	downloadCommand.Flags().StringSliceP("keys", "k", []string{}, "List of keys to download data")
	downloadCommand.Flags().StringP("targetDir", "d", "", "File to download the data to")

	err := downloadCommand.MarkFlagRequired("buckets")
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = downloadCommand.MarkFlagRequired("keys")
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = downloadCommand.MarkFlagRequired("targetDir")
	if err != nil {
		log.Fatalln(err.Error())
	}

}
