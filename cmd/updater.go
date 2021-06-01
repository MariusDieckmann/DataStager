package cmd

import (
	"context"
	"log"
	"os"

	"github.com/MariusDieckmann/DataStager/updater"
	api "github.com/ag-computational-bio/bakta-web-api-go/bakta/web/api/proto/v1"

	"github.com/spf13/cobra"
)

// uploadCommand represents the uploadCommand command
var updateCommand = &cobra.Command{
	Use:   "update",
	Short: "Updates a given job",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		runUpdate()
	},
}

func init() {
	rootCmd.AddCommand(updateCommand)
}

func runUpdate() {
	grpcUpdater := updater.InitGrpcUpdater()

	ctx := context.TODO()

	jobID := os.Getenv("JobID")
	if jobID == "" {
		log.Fatalln("No endpoint provided, please use JobID env var")
	}

	updateRequest := api.UpdateStatusRequest{
		JobID: jobID,
	}

	_, err := grpcUpdater.UpdaterClient.UpdateStatus(ctx, &updateRequest)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
