package main

import (
	"os"

	"web-app-tool/src/commands"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "QC WEB-APP-TOOL"
	app.Usage = "build and deploy web app"
	app.Version = "1.0.0"

	app.Commands = []cli.Command{
		commands.Build,
		commands.Deploy,
	}

	app.Run(os.Args)
}
