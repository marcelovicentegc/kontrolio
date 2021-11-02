package cli

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/urfave/cli/v2"
)

func Kontrolio() {

	cli.VersionPrinter = func(ctx *cli.Context) {
		fmt.Printf("version=%s\n", ctx.App.Version)
	}

	app := &cli.App{
		Name:    "kontrolio",
		Usage:   "Your cli time clock, clock card machine, punch clock or time recorder",
		Version: "0.0.46",

		Commands: []*cli.Command{
			{
				Name:    "punch",
				Aliases: []string{"p"},
				Usage:   "Punch your clock",
				Action: func(ctx *cli.Context) error {
					punch()
					status(false)
					return nil
				},
			},
			{
				Name:    "status",
				Aliases: []string{"s"},
				Usage:   "Check how many hours you have worked today",
				Action: func(ctx *cli.Context) error {
					status(true)
					return nil
				},
			},
			{
				Name:    "logs",
				Aliases: []string{"l"},
				Usage:   "Navigate through all your records",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "tail", Aliases: []string{"t"}},
					&cli.BoolFlag{Name: "graph", Aliases: []string{"g"}},
					&cli.BoolFlag{Name: "table"},
				},
				Action: func(ctx *cli.Context) error {
					tail := ctx.String("tail")
					graph := ctx.Bool("graph")
					table := ctx.Bool("table")

					if graph {
						plotGraph()
					} else if table {
						plotTable()
					} else {
						logs(&tail)
					}
					return nil
				},
			},
			{
				Name:    "configure",
				Aliases: []string{"config", "c"},
				Usage:   "Configure Kontrolio",
				Action: func(ctx *cli.Context) error {
					configure()
					return nil
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
