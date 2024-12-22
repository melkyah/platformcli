package main

import (
	"context"
	"log"
	"os"

	"github.com/melkyah/platformcli/commands/platform"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Usage:    "A CLI tool to spin up and manage infrastructure platforms",
		Version:  "v0.0.1",
		Commands: platform.Commands,
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
