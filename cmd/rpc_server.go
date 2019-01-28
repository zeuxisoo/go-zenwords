package cmd

import (
	"fmt"
	"net"
	"log"

	"google.golang.org/grpc"
	"github.com/urfave/cli"

	"github.com/zeuxisoo/go-zenwords/pkg/keywords"
	"github.com/zeuxisoo/go-zenwords/rpc/proto"
	"github.com/zeuxisoo/go-zenwords/rpc/servers"
)

var rpcServer = cli.Command{
	Name: "server",
	Usage: "Start RPC server",
	Description: "Run zenWords gRPC server for check service",
	Action: runRPCServer,
	Flags: []cli.Flag{
		stringFlag("address, a", "127.0.0.1", "Default host to run service"),
		stringFlag("port, p", "50051", "Default port number to run service"),
	},
}

func runRPCServer(c *cli.Context) error {
	log.Println("Reading the words.txt")

	keywords.NewKeywords("words.txt")

	log.Printf("--> size: %d\n", len(keywords.Words))

	address         := c.String("address")
	port            := c.String("port")
	addressWithPort := fmt.Sprintf("%s:%s", address, port)

	listener, err := net.Listen("tcp", addressWithPort)
	if err != nil {
		log.Fatalf("Fail to listen on %s: %v", addressWithPort, err)
	}

	server := grpc.NewServer()

	proto.RegisterContentServiceServer(server, &servers.ContentServiceServer{})

	log.Println("Starting RPC server")
	log.Printf("--> listen: %s:%s\n", address, port)

	server.Serve(listener)

	return nil
}
