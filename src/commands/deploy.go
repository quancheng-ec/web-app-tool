package commands

import (
	"log"
	"os"
	"path/filepath"

	"github.com/kelseyhightower/envconfig"
	"github.com/urfave/cli"
	"web-app-tool/src/ossService"
	"strings"
	"path"
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

		_, statErr := os.Stat(distPath)

		if statErr != nil {
			log.Fatal(statErr)
		}

		ossService.Init(&config)

		fmt.Println("find files in dist folder: ")

		err := filepath.Walk(distPath, func(filePath string, info os.FileInfo, err error) error {

			if !info.IsDir() && info.Name() != "index.html" {
				f, err := os.Open(filePath)
				if err != nil {
					return err
				}

				objectPath := strings.Replace(filePath, path.Clean(distPath), "", 1)

				defer f.Close()

				uploadErr := ossService.UploadToBucket(objectPath, f)

				if uploadErr != nil {
					return err
				}
			}
			return nil
		})

		if err != nil {
			log.Fatalln("upload failed for: %v", err.Error())
		}



		fmt.Println("dist files deployed")

		return nil
	},
}
