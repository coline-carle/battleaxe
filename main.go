package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/wow-sweetlie/battleaxe/battle"

	"github.com/urfave/cli"
)

func buildQueryMap(c *cli.Context) map[string]string {
	queryMap := make(map[string]string)

	locale := c.String("locale")
	if locale != "" {
		queryMap["locale"] = locale
	}

	fields := c.StringSlice("fields")
	if len(fields) > 0 {
		queryMap["fields"] = strings.Join(fields, ",")
	}

	return queryMap
}

func action(c *cli.Context) error {
	if c.NArg() != 1 {
		return cli.NewExitError("invalid number of arguments", 1)
	}

	inURL := c.Args().First()
	queryMap := buildQueryMap(c)

	outURL, err := battle.ParseURL(inURL, queryMap)

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
