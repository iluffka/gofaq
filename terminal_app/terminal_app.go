/**
*template for any app
*/

package main

import (
	"os"
	"gopkg.in/urfave/cli.v1"
	"fmt"

	"gofaq/terminal_app/config"
	"gofaq/terminal_app/helper"
	"gofaq/terminal_app/model"
)

func main() {

	app := cli.NewApp()
	app.Name = "Terminal app"
	app.Version = "0.0.1"
	app.Usage = "For all"
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: `config, c`, Usage: `path to configuration file`,}, // ./terminal_app -c path/to/config check
	}
	app.Commands = []cli.Command{
		{
			Name:        "grepos",
			Usage:       "Get all repos for user",
			Description: "Just for Description",
			ArgsUsage : "",
			Flags: []cli.Flag{
				cli.StringFlag{Name: `name, n`, Usage: `Name of user (github)`,},
			},
			Action: func(c *cli.Context) error {
				path := c.GlobalString("config")
				conf := config.Settings{}
				helper.SetConfig(path, &conf)

				userName := c.String("name")
				gitHubResponse := []model.GitHubResponse{}
				if err := helper.GetAllRepos(conf.Project.VersionControl.URL, userName, &gitHubResponse); err != nil {
					helper.ProcessError(err, "Request fail: %s")
				}
				fmt.Println(gitHubResponse)
				return nil
			},
		},
		{
			Name:			"try",
			Usage:			"Just for Usage",
			Description:	"Just for Description",
			ArgsUsage: "",
			Action: func(c *cli.Context) error {
				fmt.Println(c.Args())

				return nil
			},
		},
		// ... more cases
	}

	app.Run(os.Args)
}
