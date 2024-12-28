package platform

import (
	"context"

	"github.com/urfave/cli/v3"
)

var externalCLIActions = []string{"kind"}

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
				Commands: []*cli.Command{
					{
						Name:    "demo",
						Aliases: []string{"d"},
						Usage:   "start a demo instance in a local or remote cluster",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							cmds.kind()
							return nil
						},
					},
				},
			},
		},
	},
}
