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
	rootCmd.AddCommand(updateCommand)
	updateCommand.Flags().StringP("status", "s", "", "Updated job status")
	updateCommand.Flags().StringP("jobid", "j", "", "JobID of the job to update")
	updateCommand.Flags().StringP("secret", "k", "", "Secret of the job to update")
	updateCommand.Flags().StringP("error", "e", "", "Optionally: The reported error")

	err := updateCommand.MarkFlagRequired("status")
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = updateCommand.MarkFlagRequired("jobid")
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = updateCommand.MarkFlagRequired("secret")
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func runUpdate(jobID string, secret string, statusString string, errorString string) {
	var status api.JobStatusEnum

	switch statusString {
	case "running":
		status = api.JobStatusEnum_RUNNING
	case "error":
		status = api.JobStatusEnum_ERROR
	case "finished":
		status = api.JobStatusEnum_SUCCESFULL
	default:
		log.Fatalln(fmt.Sprintf("%v is no a valid status, please use: error, finished", statusString))
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
