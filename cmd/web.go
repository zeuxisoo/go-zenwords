package cmd

import (
	"fmt"

	"github.com/urfave/cli"
	"github.com/gin-gonic/gin"

	"github.com/zeuxisoo/go-zenwords/routes"
)

// Web command for start web service
var Web = cli.Command{
	Name: "web",
	Usage: "Start web server",
	Description: "Run zenWords web server for check service",
	Action: runWeb,
	Flags: []cli.Flag{
		boolFlag("prod", "Enable production mode"),
		stringFlag("address, a", "127.0.0.1", "Default host to run service"),
		stringFlag("port, p", "3000", "Default port number to run service"),
	},
}


func runWeb(c *cli.Context) error {
	//
	address := c.String("address")
	port := c.String("port")

	if c.IsSet("prod") {
		gin.SetMode(gin.ReleaseMode)
	}

	//
	engine := gin.Default()

	engine.GET("/", routes.HomeGet)

	engine.Run(fmt.Sprintf("%s:%s", address, port))

	return nil
}
