package urfave_cli_docs

import (
	"strconv"
	"strings"

	"github.com/urfave/cli/v2"
)

type (
	App struct {
		AppPath                     string
		Name                        string
		Usage, UsageText, ArgsUsage string
		Description                 string
		GlobalFlags                 Flags
		Commands                    []Command
	}

	Command struct {
		AppPath                     string
		Name                        string
		Aliases                     []string
		Usage, UsageText, ArgsUsage string
		Description                 string
		Category                    string
		Flags                       Flags
		SubCommands                 []Command
		Level                       uint
	}

	Flag struct {
		Name       string
		Aliases    []string
		Usage      string
		TakesValue bool
		Default    string
		EnvVars    []string
	}

	Flags []Flag
)

// NewApp creates a new App struct.
func NewApp(app *cli.App, appPath string) App {
	return App{
		AppPath:     appPath,
		Name:        app.Name,
		Description: PrepareMultilineString(app.Description),
		Usage:       PrepareMultilineString(app.Usage),
		UsageText:   PrepareMultilineString(app.UsageText),
		ArgsUsage:   PrepareMultilineString(app.ArgsUsage),
		GlobalFlags: PrepareFlags(app.VisibleFlags()),
		Commands:    PrepareCommands(app.VisibleCommands(), appPath, "", 0),
	}
}

// PrepareCommands converts CLI commands into a structs for the rendering.
func PrepareCommands(commands []*cli.Command, appPath, parentCommandName string, level uint) []Command {
	var result = make([]Command, 0, len(commands))

	for _, cmd := range commands {
		var command = Command{
			AppPath:     appPath,
			Name:        strings.TrimSpace(strings.Join([]string{parentCommandName, cmd.Name}, " ")),
			Aliases:     cmd.Aliases,
			Usage:       PrepareMultilineString(cmd.Usage),
			UsageText:   PrepareMultilineString(cmd.UsageText),
			ArgsUsage:   PrepareMultilineString(cmd.ArgsUsage),
			Description: PrepareMultilineString(cmd.Description),
			Category:    cmd.Category,
			Flags:       PrepareFlags(cmd.VisibleFlags()),
			SubCommands: PrepareCommands( // note: recursive call!
				cmd.Subcommands,
				appPath,
				strings.Join([]string{parentCommandName, cmd.Name}, " "),
				level+1,
			),
			Level: level,
		}

		result = append(result, command)
	}

	return result
}

// PrepareFlags converts CLI flags into a structs for the rendering.
func PrepareFlags(flags []cli.Flag) Flags {
	var result = make(Flags, 0, len(flags))

	for _, appFlag := range flags {
		flag, ok := appFlag.(cli.DocGenerationFlag)
		if !ok {
			continue
		}

		var f = Flag{
			Usage:      PrepareMultilineString(flag.GetUsage()),
			EnvVars:    flag.GetEnvVars(),
			TakesValue: flag.TakesValue(),
			Default:    flag.GetValue(),
		}

		if boolFlag, isBool := appFlag.(*cli.BoolFlag); isBool {
			f.Default = strconv.FormatBool(boolFlag.Value)
		}

		for i, name := range flag.Names() {
			name = strings.TrimSpace(name)

			if i == 0 {
				f.Name = "--" + name

				continue
			}

			if len(name) > 1 {
				name = "--" + name
			} else {
				name = "-" + name
			}

			f.Aliases = append(f.Aliases, name)
		}

		result = append(result, f)
	}

	return result
}

// PrepareMultilineString prepares a string (removes line breaks).
func PrepareMultilineString(s string) string {
	return strings.TrimRight(
		strings.TrimSpace(
			strings.ReplaceAll(s, "\n", " "),
		),
		".\r\n\t",
	)
}
