package cmd

import (
	"log"
	"fmt"

	"github.com/urfave/cli"
	"github.com/gin-gonic/gin"

	"github.com/zeuxisoo/go-zenwords/pkg/keywords"
	"github.com/zeuxisoo/go-zenwords/routes/home"
	"github.com/zeuxisoo/go-zenwords/routes/api"
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
	log.Println("Reading the words.txt")

	keywords.NewKeywords("words.txt")

	log.Printf("--> size: %d\n", len(keywords.Words))

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

	engine.
		Group("/api").
		POST("/search", api.SearchPost)

	//
	if isProductionMode == true {
		log.Println("Starting web server")
		log.Println("--> mode  : Production")
		log.Printf("--> listen: %s:%s\n", address, port)
	}

	//
	engine.Run(fmt.Sprintf("%s:%s", address, port))

	return nil
}
