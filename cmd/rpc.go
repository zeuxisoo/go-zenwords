package cmd

import (
	"fmt"
	"net"
	"log"

	"github.com/urfave/cli"
	"google.golang.org/grpc"

	"github.com/zeuxisoo/go-zenwords/rpc/protos"
	"github.com/zeuxisoo/go-zenwords/rpc/servers"
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



func runRPC(c *cli.Context) error {
	addressWithPort := fmt.Sprintf("%s:%s", c.String("address"), c.String("port"))

	listener, err := net.Listen("tcp", addressWithPort)
	if err != nil {
		log.Fatalf("Fail to listen on %s: %v", addressWithPort, err)
	}

	server := grpc.NewServer()

	protos.RegisterContentServiceServer(server, &servers.ContentServiceServer{})

	server.Serve(listener)

	return nil
}
