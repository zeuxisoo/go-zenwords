package cmd

import (
	"fmt"

	"github.com/urfave/cli"
)

var rpcClient = cli.Command{
	Name: "client",
	Usage: "Start RPC client",
	Description: "Run zenWords gRPC client for check service",
	Action: runRPCClient,
	Flags: []cli.Flag{
		stringFlag("address, a", "127.0.0.1", "Default host to connect the server"),
		stringFlag("port, p", "50051", "Default port for connect to the server"),
	},
}

func runRPCClient(c *cli.Context) error {
	fmt.Println("RPC Client")

	return nil
}
