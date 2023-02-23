package markdown_test

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"

	"gh.tarampamp.am/urfave-cli-docs/markdown"
)

func TestReplaceBetween(t *testing.T) {
	var src = "<foo>123</foo>"

	var out, err = markdown.ReplaceBetween("<foo>", "</foo>", src, "456")

	assert.NoError(t, err)
	assert.Equal(t, "<foo>\n456\n</foo>", out)
}

//go:embed markdown_defaults_test.md
var markdownDefaults string

func TestRenderWithAllDefaults(t *testing.T) {
	var app = &cli.App{
		Name:        "docker",
		Description: "A self-sufficient runtime for containers",
		Usage:       "Docker usage goes here",
		UsageText:   "myapp does awesome things",
		ArgsUsage:   "<any-arguments>",
		Commands: []*cli.Command{
			{
				Name:        "images",
				Description: "Manage images",
				ArgsUsage:   "<any-args>",
				Subcommands: []*cli.Command{
					{
						Name:  "ls",
						Usage: "List images",
						Flags: []cli.Flag{
							&cli.BoolFlag{
								Name:    "all",
								Aliases: []string{"a"},
								Usage:   "Show all images (default hides intermediate images)",
								Value:   true,
							},
							&cli.StringFlag{
								Name:    "filter",
								Usage:   "Filter output based on conditions provided",
								Aliases: []string{"f"},
								EnvVars: []string{"DOCKER_FILTER"},
							},
						},
					},
					{
						Name:      "rm",
						Usage:     "Remove one or more images",
						ArgsUsage: "IMAGE [IMAGE...]",
						Flags: []cli.Flag{
							&cli.BoolFlag{
								Name:    "force",
								Aliases: []string{"f"},
								Usage:   "Force removal of the image",
							},
							&cli.BoolFlag{
								Name:  "no-prune",
								Usage: "Do not delete untagged parents",
							},
						},
						Subcommands: []*cli.Command{
							{
								Name:        "test",
								Description: "Just for a test",
								Aliases:     []string{"t", "tst"},
							},
						},
					},
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "string-flag",
						Usage:   "string flag",
						Value:   "default",
						EnvVars: []string{"STRING_FLAG_ENV_VAR"},
					},
				},
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "context",
				Usage:   "Name of the context to use to connect to the daemon",
				Aliases: []string{"c"},
				EnvVars: []string{"DOCKER_CONTEXT"},
			},
			&cli.BoolFlag{
				Name:    "debug",
				Aliases: []string{"D"},
				Usage:   "Enable debug mode",
				Value:   false,
			},
			&cli.IntFlag{
				Name:    "log-level",
				Aliases: []string{"l"},
				Usage:   "Set the logging level",
				Value:   2,
			},
		},
	}

	md, err := markdown.Render(app)
	assert.NoError(t, err)

	// t.Log(md)

	assert.Equal(t, strings.TrimRight(markdownDefaults, "\n"), md)
}

func TestRenderWithOptions(t *testing.T) {
	var app = &cli.App{
		Name:        "docker",
		Description: "A self-sufficient runtime for containers",
		Usage:       "Docker usage goes here",
		UsageText:   "myapp does awesome things",
		ArgsUsage:   "<any-arguments>",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "context",
				Usage:   "Name of the context to use to connect to the daemon",
				Aliases: []string{"c"},
				EnvVars: []string{"DOCKER_CONTEXT"},
			},
		},
	}

	const tpl = `{{ .Name }} // {{ .Description }} // {{ (index .GlobalFlags 0).Name }}`

	md, err := markdown.Render(app, markdown.WithAppPath("docker"), markdown.WithTemplate(tpl))
	assert.NoError(t, err)

	assert.Equal(t, "docker // A self-sufficient runtime for containers // --context", md)
}
