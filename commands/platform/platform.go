package platform

import (
	"context"
	"log"
	"slices"

	"github.com/urfave/cli/v3"

	"github.com/melkyah/platformcli/pkg/cmds"
	"github.com/melkyah/platformcli/pkg/helm"
)

var Commands = []*cli.Command{
	{
		Name:    "platform",
		Aliases: []string{"p"},
		Usage:   "manage platform instances",
		Commands: []*cli.Command{
			{
				Name:    "install",
				Aliases: []string{"i"},
				Usage:   "bootstraps a platform instance",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "demo",
						Aliases: []string{"d"},
						Usage:   "install example demo version using Kind. Requires Podman or Docker daemon running",
					},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					Install(ctx, cmd)
					return nil
				},
			},
			{
				Name:    "uninstall",
				Aliases: []string{"u"},
				Usage:   "removes a platform instance",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "demo",
						Aliases: []string{"d"},
						Usage:   "uninstall example demo version using Kind",
					},
				}, Action: func(ctx context.Context, cmd *cli.Command) error {
					Uninstall(ctx, cmd)
					return nil
				},
			},
		},
	},
}

func Install(ctx context.Context, cmd *cli.Command) error {
	// fmt.Printf(strings.Join(cmd.FlagNames(), ", "))
	if len(cmd.FlagNames()) == 0 {
		log.Fatal("ERROR: Install command currently only support runing with --demo flag")
	} else if slices.Contains(cmd.FlagNames(), "demo") {
		args := []string{
			"create",
			"cluster",
			"--name",
			"platformcli",
		}
		cmds.Kind(args)
		return nil
	}
	return nil
}

func Uninstall(ctx context.Context, cmd *cli.Command) error {
	// fmt.Printf(strings.Join(cmd.FlagNames(), ", "))
	if len(cmd.FlagNames()) == 0 {
		log.Fatal("ERROR: Uninstall command currently only support runing with --demo flag")
	} else if slices.Contains(cmd.FlagNames(), "demo") {
		InstallDemo()
	}
	return nil
}

func InstallDemo() {
	args := []string{
		"delete",
		"cluster",
		"--name",
		"platformcli",
	}
	cmds.Kind(args)
	helm.Test()
}
