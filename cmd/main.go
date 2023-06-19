package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/tes-edot/cmd/handler"
	charpb "github.com/tes-edot/proto"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	handler := &handler.Handler{}
	charpb.RegisterChartServer(s, handler)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", 8080))

	if err != nil {
		os.Exit(1)
	}
	log.Println("Starting RPC server at", 8080)
	if err != nil {
		log.Fatalf("could not listen to %v: %v", 8080, err)
	}
	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve %w", err)
		os.Exit(1)
	}
}
