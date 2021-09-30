package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

// Build time values.
//
// We will set this value via `go build -ldflags "-X main.Version"`
var (
	Version string
)

var app = cli.App{
	Name:        "btp",
	Description: "BeyondTP is a neutral data migration service",
	Version:     Version,
	Commands: []*cli.Command{
		agentCmd,
	},
}

func main() {
	err := app.Run(os.Args)
	if err != nil {
		// FIXME: we need to respect platform style later.
		os.Exit(1)
	}
}

func userConfigDir() string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		panic("$HOME is not specified")
	}
	return configDir
}
