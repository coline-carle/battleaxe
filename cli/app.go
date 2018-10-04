package cli

import (
	"fmt"
	"log"
	"os"
	"golang.org/x/oauth2/clientcredentials"
	"golang.org/x/oauth2"

	"github.com/coline-carle/battleaxe/battle"
)

// Version of the app
const Version = "0.0.2"

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


	var clientID string
	if flags.clientID == "" {
		clientID = os.Getenv("BLIZZARD_CLIENT_ID")
	} else {
		clientID = flags.clientID 
	}

	var clientSecret string
	if flags.clientSecret == "" {
		clientSecret = os.Getenv("BLIZZARD_CLIENT_SECRET")
	} else {
		clientSecret = flags.clientSecret 
	}

	if clientID == "" {
		logger.Println("clientID can't be empty use --clientid flag or BLIZZARD_CLIENT_ID env variable")
	}

	if clientSecret == "" {
		logger.Println("clientSecret can't be empty use --clientid flag or BLIZZARD_CLIENT_SECRET env variable")
	}

	// if client secret is  not set try to fetch it from env
	if flags.clientID == "" {
		flags.clientSecret = os.Getenv("BLIZZARD_CLIENT_ID")
	}

	queryMap := buildQueryMap(flags)

	blizzOauth := &clientcredentials.Config{
		ClientID:     clientID,
    ClientSecret: clientSecret,
		TokenURL: "https://us.battle.net/oauth/token",
	}

	client := blizzOauth.Client(oauth2.NoContext)


	url, err := battle.ParseURL(inURL, queryMap, game)

	if err != nil {
		logger.Fatal(err)
	}

	if flags.dry {
		fmt.Println(url)
		os.Exit(0)
	}

	resp, err := client.Get(url)
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
