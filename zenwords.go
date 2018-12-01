package main

import (
	"os"

	"github.com/urfave/cli"

	"github.com/zeuxisoo/go-zenwords/cmd"
)

// AppVer is A application version
const AppVer = "0.1.0"

func main() {
	app := cli.NewApp()
	app.Name = "ZenWords"
	app.Usage = "A zen words check service"
	app.Version = AppVer
	app.Commands = []cli.Command{
		cmd.Web,
	}
	app.Run(os.Args)
}
