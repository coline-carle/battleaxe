package cli

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/wow-sweetlie/battleaxe/battle"
)

// Version of the app
const Version = "0.0.1"

var logger *log.Logger

// AppName : determine app name bbased on Game value
func AppName(game battle.Game) string {
	switch game {
	case battle.WoW:
		return "wowaxe"
	case battle.D3:
		return "daxe"
	case battle.SC2:
		return "scaxe"

	default:
		return "battleaxe"
	}
}

func init() {
	logger = log.New(os.Stderr, "", 0)
}

func buildQueryMap(f *appFlags) map[string]string {
	queryMap := make(map[string]string)

	if f.locale != "" {
		queryMap["locale"] = f.locale
	}

	if f.fields != "" {
		queryMap["fields"] = f.fields
	}

	if f.apikey != "" {
		queryMap["apikey"] = f.apikey
	}

	return queryMap
}

// Run the app
func Run(game battle.Game, args []string) {
	flags, inURL, err := parseCommand(args[1:])
	if err != nil {
		logger.Println(err)
		err = PrintHelp(AppName(game))
		if err != nil {
			logger.Println(err)
		}
		os.Exit(1)
	}

	if flags.version {
		PrintVersion()
		os.Exit(0)
	}

	if flags.help {
		err = PrintHelp(AppName(game))
		if err != nil {
			logger.Fatal(err)
		}
		os.Exit(0)
	}

	// if apikey not set try to fetch it from env
	if flags.apikey == "" {
		flags.apikey = os.Getenv("BATTLENET_CLIENT_ID")
	}

	queryMap := buildQueryMap(flags)

	url, err := battle.ParseURL(inURL, queryMap, game)

	if err != nil {
		logger.Fatal(err)
	}

	if flags.dry {
		fmt.Println(url)
		os.Exit(0)
	}

	resp, err := http.Get(url)
	if err != nil {
		logger.Fatal(err)
	}

	if flags.head {
		PrintHeader(resp)
		os.Exit(0)
	}

	err = PrintBody(resp, flags.human)
	if err != nil {
		logger.Fatal(err)
	}
	os.Exit(0)
}
