package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/MariusDieckmann/DataStager/updater"

	"github.com/ag-computational-bio/bakta-web-api/go/api"

	"github.com/spf13/cobra"
)

// uploadCommand represents the uploadCommand command
var updateCommand = &cobra.Command{
	Use:   "update",
	Short: "Updates a given job",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		status, err := cmd.Flags().GetString("status")
		if err != nil {
			log.Fatalln(err.Error())
		}

		JobID, err := cmd.Flags().GetString("jobid")
		if err != nil {
			log.Fatalln(err.Error())
		}

		secret, err := cmd.Flags().GetString("secret")
		if err != nil {
			log.Fatalln(err.Error())
		}

		error, err := cmd.Flags().GetString("error")
		if err != nil {
			log.Fatalln(err.Error())
		}

		runUpdate(JobID, secret, status, error)

	},
}

func init() {
	rootCmd.AddCommand(uploadCommand)
	uploadCommand.Flags().StringP("status", "s", "", "Updated job status")
	uploadCommand.Flags().StringP("jobid", "j", "", "JobID of the job to update")
	uploadCommand.Flags().StringP("secret", "k", "", "Secret of the job to update")
	uploadCommand.Flags().StringP("error", "k", "", "Optionally: The reported error")

	err := uploadCommand.MarkFlagRequired("status")
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = uploadCommand.MarkFlagRequired("jobid")
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = uploadCommand.MarkFlagRequired("secret")
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func runUpdate(jobID string, secret string, statusString string, errorString string) {
	var status api.JobStatusEnum

	switch statusString {
	case "error":
		status = api.JobStatusEnum_ERROR
	case "finished":
		status = api.JobStatusEnum_RUNNING
	default:
		log.Fatalln(fmt.Sprintf("%v is no a valid status, please use: error, finished"))
	}

	grpcUpdater := updater.InitGrpcUpdater()

	ctx := context.TODO()

	updateRequest := api.UpdateStatusRequest{
		JobID:  jobID,
		Secret: secret,
		Status: status,
		Error:  errorString,
	}

	_, err := grpcUpdater.UpdaterClient.UpdateStatus(ctx, &updateRequest)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
