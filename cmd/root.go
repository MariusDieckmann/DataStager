package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	//CustomS3Endpoint Variables to store and optional custom s3 endpoint
	CustomS3Endpoint string
	userLicense      string

	rootCmd = &cobra.Command{
		Use:   "DataStager",
		Short: "A simple CLI tool to stage data from and to S3 object storage in Kubernetes",
		Long:  ``,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

//Basic configuration options
func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	rootCmd.PersistentFlags().StringVarP(&CustomS3Endpoint, "s3endpoint", "e", "", "S3 endpoint for a custom s3 deployment, e.g.: s3.example.com")
	viper.SetDefault("license", "apache")

}

func initConfig() {

}
