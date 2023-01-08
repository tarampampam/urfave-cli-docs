package markdown_test

import (
	"testing"

	"github.com/urfave/cli/v2"

	"github.com/tarampampam/urfave-cli-docs/markdown"
)

func TestToMarkdown(t *testing.T) {
	var app = &cli.App{
		Name:      "myapp",
		Usage:     "does awesome things",
		UsageText: "myapp does awesome things",
		Commands: []*cli.Command{
			{
				Name:  "cmd",
				Usage: "does awesome things",
				Subcommands: []*cli.Command{
					{
						Name:  "subcmd",
						Usage: "does awesome things",
					},
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "string-flag",
						Usage:   "string flag",
						Value:   "default",
						EnvVars: []string{"STRING_FLAG_ENV_VAR"},
					},
					&cli.IntFlag{
						Name:    "int-flag",
						Usage:   "int flag",
						Value:   123,
						EnvVars: []string{"INT_FLAG_ENV_VAR"},
					},
				},
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "app-string-flag",
				Usage:   "app string flag",
				Aliases: []string{"a"},
				Value:   "default",
				EnvVars: []string{"STRING_FLAG_ENV_VAR"},
			},
			&cli.IntFlag{
				Name:    "app-int-flag",
				Usage:   "app int flag",
				Value:   123,
				EnvVars: []string{"INT_FLAG_ENV_VAR"},
			},
			&cli.IntFlag{
				Name:  "l",
				Usage: "short app int flag",
				Value: 321,
			},
		},
	}

	t.Log(markdown.ToMarkdown(app))
}
