package cli

import (
	"os"
	"strings"
	"text/template"
)

const (
	usageLocale       = "game locale example: en_US"
	usageClientID     = "blizzard API client id, environnement variable BLIZZARD_CLIENT_ID is used by default if unset"
	usageClientSecret = "blizzard API client secret, environnement variable BLIZZARD_CLIENT_SECRET is used by default if unset"
	usageHuman        = "humanize response with color and indentation"
	usagePretty       = usageHuman + " (same as --human)"
	usageHead         = "print headers instead of body"
	usageVersion      = "show version"
	usageDry          = "print the url that would be request instead of fetching it"
	usageHelp         = "show this help"
	usageFields       = "set optional fields for the requested endpoint"
)

var appHelpTemplate = `
USAGE:
  {{.Name}} {{.Usage}}

DESCRIPTION:
	examples of urls: {{range .Examples }}
	  {{.}}{{end}}

	Golden rule: Flags that modify query options, always take precedence
	over the query option of the url. In case of multiple definition of the same
	flag: The rightest flag win.

QUERY OPTIONS: {{range .QueryFlags}}
  {{join .Names ", "}}{{"\t"}}{{.Usage}}{{end}}

GLOBAL OPTIONS: {{range .GeneralFlags}}
  {{join .Names ", "}}{{"\t"}}{{.Usage}}{{end}}

SEE ALSO:
	{{join .Apps ", "}}

VERSION:
  {{.Version}}
`

// FlagHelp : context for flag help
type FlagHelp struct {
	Names []string
	Usage string
}

// Description of the app

// AppHelp : base struct for template parsing
type AppHelp struct {
	Name         string
	Usage        string
	Apps         []string
	Examples     []string
	QueryFlags   []FlagHelp
	GeneralFlags []FlagHelp
	Version      string
}

var queryFlags = []FlagHelp{
	{
		[]string{"client", "K"},
		usageClientID,
	},
	{
		[]string{"secret", "S"},
		usageClientSecret,
	},
	{
		[]string{"fields", "F"},
		usageFields,
	},
	{
		[]string{"locale", "L"},
		usageLocale,
	},
}

var generalFlags = []FlagHelp{
	{
		[]string{"head", "I"},
		usageHead,
	},
	{
		[]string{"human", "C"},
		usageHuman,
	},
	{
		[]string{"pretty"},
		usageHuman + " (same as human)",
	},
	{
		[]string{"version", "V"},
		usageVersion,
	},
	{
		[]string{"dry", "D"},
		usageDry,
	},
	{
		[]string{"help", "usage"},
		usageHelp,
	},
}

var urlExamples = map[string][]string{
	"battleaxe": {
		"https://us.api.blizzard.com/wow/achievement/2144?locale=en_US&apikey=APIKEY",
		"us://wow/achievement/2144",
		"wow://achievement/2144",
	},
	"wowaxe": {
		"https://us.api.blizzard.com/wow/achievement/2144?locale=en_US&apikey=APIKEY",
		"us://achievement/2144",
		"achievement://2144",
	},
	"scaxe": {
		"https://us.api.blizzard.com/sc2/ladder/194163?locale=en_US&apikey=APIKEY",
		"us://ladder/2144",
		"ladder://2144",
	},
	"daxe": {
		"https://us.api.blizzard.com/d3/data/follower/templar?locale=en_US&apikey=APIKEY",
		"us://data/follower/templar",
		"data://follower/templar",
	},
}

// PrintHelp : ...
func PrintHelp() error {
	funcMap := template.FuncMap{
		"join": strings.Join,
	}

	appName := os.Args[0]
	help := &AppHelp{
		Name:         appName,
		GeneralFlags: generalFlags,
		QueryFlags:   queryFlags,
		Apps:         []string{"battleaxe", "wowaxe", "scaxe", "daxe"},
		Examples:     urlExamples[appName],
		Usage:        "[OPTIONS] scheme://path/to/resource [MORE OPTIONS]",
		Version:      Version,
	}

	tmpl, err := template.New("help").Funcs(funcMap).Parse(appHelpTemplate)
	if err != nil {
		return err
	}

	err = tmpl.Execute(os.Stdout, help)
	if err != nil {
		return err
	}

	return nil
}
