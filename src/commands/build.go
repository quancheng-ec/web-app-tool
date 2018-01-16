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

var buildFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "tag,t",
		Value: time.Now().Format("20060102150405"),
		Usage: "tag for build",
	},
	cli.StringFlag{
		Name:  "cdn-domain",
		Value: "//s-cdn.qccost.com",
		Usage: "cdn domain, no schema",
	},
	cli.StringFlag{
		Name:  "name,n",
		Usage: "project name",
	},
}

var Build = cli.Command{
	Name:      "build",
	Aliases:   []string{"b"},
	Usage:     "build web-app",
	Flags:     buildFlags,
	ArgsUsage: "project-path",
	Action: func(c *cli.Context) error {

		projectPath := c.Args().First()

		if len(projectPath) == 0 {
			projectPath = "."
		}

		fmt.Println("build path found: ", projectPath)

		packageInfo, _ := utils.ReadPackageInfo(projectPath)

		appName := packageInfo.Name

		if len(c.String("name")) != 0 {
			appName = c.String("name")
		}

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

		fmt.Printf("building %s starts ...", appName)

		buildEnv := make(map[string]string)

		buildEnv["BUILD_TAG"] = c.String("tag")
		buildEnv["CDN_PATH"] = c.String("cdn-domain")
		buildEnv["APP_NAME"] = appName

		for k, v := range buildEnv {
			buildSession.SetEnv(k, v)
		}

		fmt.Println("build env: ", buildEnv)

		buildSession.Command(npmPath, "run", "build").Run()

		fmt.Printf("%s has been built", appName)

		return nil
	},
}
