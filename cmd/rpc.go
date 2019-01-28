package cmd

import (
	"github.com/urfave/cli"
)

// RPC command for start RPC service
var RPC = cli.Command{
	Name: "rpc",
	Usage: "Start RPC service",
	Description: "Run zenWords gRPC server / client for check service",
	Subcommands: []cli.Command{
		rpcServer,
		rpcClient,
	},
}
