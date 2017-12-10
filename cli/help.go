package cli

import (
	"os"
	"strings"
	"text/template"
)

var appHelpTemplate = `NAME:
  {{.Name}}

USAGE:
  {{.Name}} {{.Usage}}

DESCRIPTION:
	examples of urls: {{range .Examples }}
	  {{.}}{{end}}

	Golden rule: Flags that modify query options, always take precedance
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
		[]string{"apikey", "k"},
		"your personal api key variable env BATTLENET_CLIENT_ID is used by default if set",
	},
	{
		[]string{"fields", "f"},
		"set fields for endpoint that accept ones",
	},
	{
		[]string{"locale", "l"},
		"game locale ex: en_US",
	},
}

var generalFlags = []FlagHelp{
	{
		[]string{"head", "I"},
		"print headers instead of body",
	},
	{
		[]string{"color", "C"},
		"humanize response with color and indentation",
	},
	{
		[]string{"version", "V"},
		"show version",
	},
	{
		[]string{"dry", "D"},
		"only print the url that would be requested",
	},
	{
		[]string{"help", "usage"},
		"show this help",
	},
}

var urlExemples = map[string][]string{
	"battleaxe": []string{
		"https://us.api.battle.net/wow/achievement/2144?locale=en_US&apikey=APIKEY",
		"us://wow/achievement/2144",
		"wow://achievement/2144",
	},
	"wowaxe": []string{
		"https://us.api.battle.net/wow/achievement/2144?locale=en_US&apikey=APIKEY",
		"us://achievement/2144",
		"achievement://2144",
	},
	"scaxe": []string{
		"https://us.api.battle.net/sc2/ladder/194163?locale=en_US&apikey=APIKEY",
		"us://ladder/2144",
		"ladder://2144",
	},
}

// PrintHelp : ...
func PrintHelp(appname string) error {
	funcMap := template.FuncMap{
		"join": strings.Join,
	}
	help := &AppHelp{
		Name:         appname,
		GeneralFlags: generalFlags,
		QueryFlags:   queryFlags,
		Apps:         []string{"battleaxe", "wowaxe", "scaxe", "daxe"},
		Examples:     urlExemples[appname],
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
