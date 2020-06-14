package main

import (
	"github.com/urfave/cli/v2"
)

var version string

func main() {
	app := &cli.App{
		Name:        "serveup",
		Usage:       "CLI tools over HTTP",
		Description: "Proxies HTTP calls to CLI apps",
		Version:     version,
		Commands: []*cli.Command{
			{
				Name:    "run",
				Aliases: []string{"c"},
				Usage:   "complete a task on the list",
				Action:  func(c *cli.Context) error {
					return nil
				},
			},
		},
	}
}