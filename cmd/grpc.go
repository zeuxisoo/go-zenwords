package cmd

import (
	"fmt"

	"github.com/urfave/cli"
)

// Rpc command for start RPC service
var Rpc = cli.Command{
	Name: "rpc",
	Usage: "Start RPC server",
	Description: "Run zenWords gRPC server for check service",
	Action: runRpc,
	Flags: []cli.Flag{
	},
}

func runRpc(c *cli.Context) error {
	fmt.Println("RPC command")

	return nil
}
