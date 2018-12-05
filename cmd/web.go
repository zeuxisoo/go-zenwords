package cmd

import (
	"fmt"

	"github.com/urfave/cli"
	"github.com/gin-gonic/gin"

	"github.com/zeuxisoo/go-zenwords/routes/home"
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
	isProductionMode := c.IsSet("prod")

	if isProductionMode == true {
		gin.SetMode(gin.ReleaseMode)
	}

	//
	engine := gin.Default()

	engine.GET("/", home.IndexGet)
	engine.GET("/robots.txt", home.RobotsTxtGet)

	//
	if isProductionMode == true {
		fmt.Printf("Listen on %s:%s\n", address, port)
	}

	//
	engine.Run(fmt.Sprintf("%s:%s", address, port))

	return nil
}
