package platform

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
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
				Commands: []*cli.Command{
					{
						Name:    "demo",
						Aliases: []string{"p"},
						Usage:   "start a demo instance in a local cluster",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("new task template: ", cmd.Args().First())
							return nil
						},
					},
				},
			},
		},
	},
}
