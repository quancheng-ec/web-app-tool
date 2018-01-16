package commands

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/urfave/cli"
	"github.com/codeskyblue/go-sh"
	"web-app-tool/src/utils"
	"time"
)

var Build = cli.Command{
	Name:    "build",
	Aliases: []string{"b"},
	Usage:   "build web-app",
	Action: func(c *cli.Context) error {

		projectPath := c.Args().First()

		if len(projectPath) == 0 {
			projectPath = "."
		}

		fmt.Println("build path found: ", projectPath)

		packageInfo, _ := utils.ReadPackageInfo(projectPath)

		fmt.Println("app name: ", packageInfo.Name)
		fmt.Println("app version: ", packageInfo.Version)

		npmPath, err := exec.LookPath("npm")

		fmt.Println("build env checking ...")

		if err != nil {
			log.Fatal("npm is not found")
		}

		fmt.Println("build env checked!")

		fmt.Println("installing dependencies ...")

		buildSession := sh.NewSession()

		buildSession.SetDir(projectPath)

		buildSession.ShowCMD = true

		buildSession.Command(npmPath, "install").Run()

		fmt.Println("dependencies installed ...")

		fmt.Println("building web-app ...")

		buildEnv := make(map[string]string)

		buildEnv["BUILD_TAG"] = time.Now().Format("20060102150405")
		buildEnv["CDN_PATH"] = "//s-cdn.qccost.com/" + packageInfo.Name

		for k, v := range buildEnv {
			buildSession.SetEnv(k, v)
		}

		fmt.Println("build env: ", buildEnv)

		buildSession.Command(npmPath, "run", "build").Run()

		return nil
	},
}
