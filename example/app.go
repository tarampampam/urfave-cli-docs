package example

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func NewApp() *cli.App {
	return &cli.App{
		Name:  "example",
		Usage: "Example CLI application",
		Commands: []*cli.Command{
			{
				Name:    "hello",
				Aliases: []string{"hi"},
				Usage:   "Prints hello",
				Action: func(c *cli.Context) error {
					fmt.Printf("Hello, %s!", c.String("name"))
					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "name",
						Value:   "stranger",
						Usage:   "Just a name",
						Aliases: []string{"n"},
						EnvVars: []string{"NAME"},
					},
				},
			},
		},
	}
}
