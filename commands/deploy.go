package commands

import (
	"log"
	"os"
	"path/filepath"

	"web-app-tool/ossService"

	"github.com/kelseyhightower/envconfig"
	"github.com/urfave/cli"
	"fmt"
)

var Deploy = cli.Command{
	Name:    "deploy",
	Aliases: []string{"d"},
	Usage:   "deploy built web-app",
	Action: func(c *cli.Context) error {

		var config ossService.ServiceConfig

		envErr := envconfig.Process("", &config)

		if envErr != nil {
			log.Fatal(envErr)
		}

		distPath := c.Args().First()

		if distPath == "" {
			distPath = "dist"
		}

		ossService.Init(&config)

		err := filepath.Walk(distPath, func(filePath string, info os.FileInfo, err error) error {

			if !info.IsDir() && info.Name() != "index.html" {
				f, err := os.Open(filePath)
				if err != nil {
					return err
				}

				fmt.Println(filePath)

				defer f.Close()
				//uploadErr := ossService.UploadToBucket(objectPath, f)
				//
				//if uploadErr != nil {
				//	return err
				//}
			}
			return nil
		})

		if err != nil {
			log.Fatalln("upload failed for: %v", err.Error())
		}

		return nil
	},
}
