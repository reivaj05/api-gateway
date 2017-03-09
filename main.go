package main

import (
	"fmt"
	"os"

	"github.com/reivaj05/apigateway/generator"

	"github.com/reivaj05/apigateway/cli"
	"github.com/reivaj05/apigateway/server"
)

// TODO: Validate all things needed(GOPATH is set, etc.)

const appName = "apigateway"

func main() {
	setup()
	startApp()
}

func setup() {
	startConfig()
	startLogger()
}

func startConfig() {

}

func startLogger() {

}

func startApp() {
	if err := cli.StartCLI(createCLIOptions()); err != nil {
		finishExecution("Error while starting application", map[string]interface{}{
			"error": err.Error(),
		})
	}
}

func createCLIOptions() *cli.Options {
	return &cli.Options{
		AppName:       appName,
		AppUsage:      "TODO: Set app usage",
		Commands:      createCommands(),
		DefaultAction: server.Serve,
	}
}

func createCommands() []*cli.Command {
	return []*cli.Command{
		&cli.Command{
			Name:   "start",
			Usage:  "TODO: Set start usage",
			Action: server.Serve,
		},
		&cli.Command{
			Name:   "create-service",
			Usage:  "TODO: Set create-service usage",
			Action: generator.Generate,
		},
	}
}

func finishExecution(msg string, fields map[string]interface{}) {
	fmt.Println(msg, fields)
	os.Exit(1)
}