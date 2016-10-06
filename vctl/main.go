package main

import (
	"os"

	"github.com/ruprict/vulcand/plugin/registry"
	"github.com/ruprict/vulcand/vctl/command"
)

var vulcanUrl string

func main() {
	cmd := command.NewCommand(registry.GetRegistry())
	err := cmd.Run(os.Args)
	if err != nil {
		cmd.PrintError(err)
	}
}
