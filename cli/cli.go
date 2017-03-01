package cli

import (
	"os"

	"github.com/urfave/cli"
)

type Command struct {
	Name   string
	Usage  string
	Action func(args ...interface{}) error
}

type StringFlag struct {
	*cli.StringFlag
}

type BoolFlag struct {
	*cli.BoolFlag
}

type Options struct {
	AppName       string
	AppUsage      string
	Commands      []*Command
	StringFlags   []*StringFlag
	BoolFlags     []*BoolFlag
	DefaultAction func(args ...interface{}) error
}

func StartCLI(options *Options) error {
	return createCLIApp(options).Run(os.Args)
}

func createCLIApp(options *Options) *cli.App {
	app := cli.NewApp()
	app.Name = options.AppName
	app.Usage = options.AppUsage
	app.Commands = createAppCommands(options.Commands)
	// app.Flags = createAppFlags()
	if options.DefaultAction != nil {
		app.Action = func(c *cli.Context) error {
			return options.DefaultAction()
		}
	}
	return app
}

func createAppCommands(commands []*Command) cli.Commands {
	co := cli.Commands{}
	for _, command := range commands {
		co = append(co, createCommand(command))
	}
	return co
}

func createCommand(command *Command) cli.Command {
	return cli.Command{
		Name:  command.Name,
		Usage: command.Usage,
		Action: func(c *cli.Context) error {
			return command.Action()
		},
	}
}

func createAppFlags() []*cli.Command {
	return nil
}
