package cmd

import (
	"fmt"
	"net"
	"log"
	"context"

	"github.com/urfave/cli"
	"google.golang.org/grpc"

	"github.com/zeuxisoo/go-zenwords/rpc/protos"
)

// RPC command for start RPC service
var RPC = cli.Command{
	Name: "rpc",
	Usage: "Start RPC server",
	Description: "Run zenWords gRPC server for check service",
	Action: runRPC,
	Flags: []cli.Flag{
		stringFlag("address, a", "127.0.0.1", "Default host to run service"),
		stringFlag("port, p", "50051", "Default port number to run service"),
	},
}

type contentServiceServer struct {}

func (s *contentServiceServer) Replace(context context.Context, request *protos.ContentReplaceRequest) (*protos.ContentReplaceResponse, error) {
	return &protos.ContentReplaceResponse{
		Result: "TODO: filter content action",
	}, nil
}

func runRPC(c *cli.Context) error {
	addressWithPort := fmt.Sprintf("%s:%s", c.String("address"), c.String("port"))

	listener, err := net.Listen("tcp", addressWithPort)
	if err != nil {
		log.Fatalf("Fail to listen on %s: %v", addressWithPort, err)
	}

	server := grpc.NewServer()

	protos.RegisterContentServiceServer(server, &contentServiceServer{})

	server.Serve(listener)

	return nil
}
