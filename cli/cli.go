package cli

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/marcelovicentegc/kontrolio-cli/config"
	"github.com/urfave/cli/v2"
)

func Kontrolio() {
	config.ConfigNetworkMode()

	cli.VersionPrinter = func(ctx *cli.Context) {
		fmt.Printf("version=%s\n", ctx.App.Version)
	}

	app := &cli.App{
		Name:    "kontrolio",
		Usage:   "your cli time clock, clock card machine, punch clock or time recorder",
		Version: "0.0.25",

		Commands: []*cli.Command{
			{
				Name:    "punch",
				Aliases: []string{"p"},
				Usage:   "punch your clock",
				Action: func(ctx *cli.Context) error {
					punch()
					workday(false)
					return nil
				},
			},
			{
				Name:    "status",
				Aliases: []string{"s"},
				Usage:   "check how many hours have you worked today",
				Action: func(ctx *cli.Context) error {
					workday(true)
					return nil
				},
			},
			{
				Name:    "sync",
				Usage:   "sync offline and online records",
				Action: func(ctx *cli.Context) error {
					sync()
					return nil
				},
			},
			{
				Name:    "logs",
				Aliases: []string{"l"},
				Usage:   "navigate through all your records",
				Action: func(ctx *cli.Context) error {
					logs()
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
