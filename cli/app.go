package cli

import (
	"log"
	"net/http"
	"os"

	"github.com/wow-sweetlie/battleaxe/battle"
)

const version = "0.0.1"

var logger *log.Logger

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
		logger.Fatal(err)
	}

	// if apikey not set try to fetch it from env
	if flags.apikey != "" {
		flags.apikey = os.Getenv("BATTLENET_CLIENT_ID")
	}

	queryMap := buildQueryMap(flags)

	url, err := battle.ParseURL(inURL, queryMap, game)

	if err != nil {
		logger.Fatal(err)
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
