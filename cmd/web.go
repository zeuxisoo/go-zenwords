package cmd

import (
	"fmt"

	"github.com/urfave/cli"
)

// Web command for start web service
var Web = cli.Command{
	Name: "web",
	Usage: "Start web server",
	Description: "Run zenWords web server for check service",
	Action: runWeb,
	Flags: []cli.Flag{
		stringFlag("address, a", "127.0.0.1", "Default host to run service"),
		stringFlag("port, p", "3000", "Default port number to run service"),
	},
}


func runWeb(c *cli.Context) error {
	fmt.Println("TODO: start web server")

	return nil
}
