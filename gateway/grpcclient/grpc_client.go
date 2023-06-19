package grpcclient

import (
	"fmt"
	"log"

	charpb "github.com/tes-edot/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CharService() charpb.ChartClient {
	port := "8080"

	conn, err := grpc.Dial(fmt.Sprintf("localhost:%s", port), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal("could not connect to :", port, err)
	}
	return charpb.NewChartClient(conn)

}
