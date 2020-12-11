package updater

import (
	"fmt"
	"log"
	"os"

	"github.com/ag-computational-bio/bakta-web-api/go/api"
	"google.golang.org/grpc"
)

type GrpcUpdater struct {
	UpdaterClient api.BaktaStatusUpdateClient
}

func InitGrpcUpdater() *GrpcUpdater {
	var opts []grpc.DialOption

	endpoint := os.Getenv("GRPCUpdaterEndpoint")
	if endpoint == "" {
		log.Fatalln("No endpoint provided, please use GRPCUpdaterEndpoint env var")
	}

	port := os.Getenv("GRPCUpdaterPort")
	if endpoint == "" {
		log.Fatalln("No endpoint port provided, please use GRPCUpdaterPort env var")
	}

	opts = append(opts, grpc.WithInsecure())

	serverAddr := fmt.Sprintf("%v:%v", endpoint, port)

	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	updaterClient := api.NewBaktaStatusUpdateClient(conn)

	gprcUpdater := GrpcUpdater{
		UpdaterClient: updaterClient,
	}

	return &gprcUpdater

}
