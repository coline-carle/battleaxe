package main

import (
	"fmt"
	"os"

	"github.com/wow-sweetlie/battleaxe/battle"

	"github.com/urfave/cli"
)

func action(c *cli.Context) error {
	if c.NArg() != 1 {
		return cli.NewExitError("invalid number of arguments", 1)
	}

	inURL := c.Args().First()
	outURL, err := battle.ParseURL(inURL)
	if err != nil {
		return cli.NewExitError(err, 1)
	}
	fmt.Println(outURL)
	return nil
}

func main() {

	app := cli.NewApp()
	app.Name = "battleaxe"
	app.Usage = "region://game/path"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "locale, l",
			Usage: "`locale`, ex: --locale en_US (default is non-set)",
		},
		cli.StringFlag{
			Name:   "apikey, k",
			EnvVar: "BATTLENET_CLIENT_ID",
			Usage:  "your personal `apikey`",
		},
		cli.StringSliceFlag{
			Name:  "fields, f",
			Usage: "set up `fields` for method that have ones. (override the ones present in url)",
		},
	}
	app.Action = action

	app.Run(os.Args)
}
