package main

import (
	"context"
	"fmt"
	"os"
	"slices"
	"strings"
	"testing"

	"github.com/melkyah/platformcli/commands/platform"

	"github.com/urfave/cli/v3"
)

func Test_CommandErrReturn(t *testing.T) {
	t.Parallel()
	cases := []struct {
		Name       string
		Command    string
		Subcommand []string
		WantOut    string
	}{
		{
			Name:       "unknown flag in platform install command",
			Command:    "platform",
			Subcommand: []string{"install", "--wrong"},
			WantOut:    "flag provided but not defined: -wrong",
		},
		{
			Name:       "unknown flag in platform uninstall command",
			Command:    "platform",
			Subcommand: []string{"uninstall", "--wrong"},
			WantOut:    "flag provided but not defined: -wrong",
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			cmd := &cli.Command{
				Usage:    "A CLI tool to spin up and manage infrastructure platforms",
				Version:  "v0.0.1",
				Commands: platform.Commands,
			}
			arguments := os.Args[0:1]
			arguments = slices.Concat(arguments, []string{tc.Command}, tc.Subcommand)
			t.Parallel()
			fmt.Printf("Running command %s \n", strings.Join(arguments, " "))
			err := cmd.Run(context.Background(), arguments)
			if err == nil {
				t.Errorf("Subcommand should raise an error if not called with correct params")
			} else if err.Error() != tc.WantOut {
				t.Errorf("Error output not matching.\n Expected: %s\n Received: %s", tc.WantOut, err.Error())
			} else {
				fmt.Printf("Error:\n%s\n", err.Error())
			}
		})
	}
}

func Test_CreateDemoCluster(t *testing.T) {
	cases := []struct {
		Name       string
		Command    string
		Subcommand []string
	}{
		{
			Name:       "create Kind cluster",
			Command:    "platform",
			Subcommand: []string{"install", "--demo"},
		},
		{
			Name:       "delete Kind cluster",
			Command:    "platform",
			Subcommand: []string{"uninstall", "--demo"},
		},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			cmd := &cli.Command{
				Usage:    "A CLI tool to spin up and manage infrastructure platforms",
				Version:  "v0.0.1",
				Commands: platform.Commands,
			}
			arguments := os.Args[0:1]
			arguments = slices.Concat(arguments, []string{tc.Command}, tc.Subcommand)

			t.Cleanup(func() {
				cmd.Run(context.Background(), slices.Concat(os.Args[0:1], []string{"platform", "uninstall", "--demo"}))
			})

			fmt.Printf("Running command %s \n", strings.Join(arguments, " "))
			err := cmd.Run(context.Background(), arguments)
			if err != nil {
				t.Errorf("Error running Kind operation:\n%s", err.Error())
			}
		})
	}
}
