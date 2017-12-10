package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/wow-sweetlie/battleaxe/battle"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stderr, "", 0)
}

type appFlags struct {
	locale  string
	fields  string
	apikey  string
	version bool
}

type context struct {
	queryMap map[string]string
	flags    *appFlags
	url      string
}

func firstString(a string, b string) string {
	if a == "" {
		return b
	}
	return a
}

func mergeFlags(flags *appFlags, afterFlags *appFlags) *appFlags {
	return &appFlags{
		locale:  firstString(afterFlags.locale, flags.locale),
		fields:  firstString(afterFlags.fields, flags.fields),
		apikey:  firstString(afterFlags.apikey, flags.apikey),
		version: flags.version || afterFlags.version,
	}
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

func (c *context) display(resp *http.Response) error {
	defer resp.Body.Close()

	_, err := io.Copy(os.Stdout, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func (c *context) action() {

	url, err := battle.ParseURL(c.url, c.queryMap)

	if err != nil {
		logger.Fatal(err)
	}

	resp, err := http.Get(url)
	if err != nil {
		logger.Fatal(err)
	}

	err = c.display(resp)
	if err != nil {
		logger.Fatal(err)
	}
}

func parseFlags(args []string) (*appFlags, []string, error) {
	flags := &appFlags{}
	apikeyFromEnv := os.Getenv("BATTLENET_CLIENT_ID")

	flagset := flag.NewFlagSet("battleaxe", -1)

	localeUsage := "locale"
	flagset.StringVar(&flags.locale, "locale", "", localeUsage)
	flagset.StringVar(&flags.locale, "l", "", localeUsage)

	apikeyUsage := "your personal api key"
	flagset.StringVar(&flags.apikey, "apikey", apikeyFromEnv, apikeyUsage)
	flagset.StringVar(&flags.apikey, "k", apikeyFromEnv, apikeyUsage)

	fieldsUsage := "select fields to fetch on endpoint with this option"
	flagset.StringVar(&flags.fields, "fields", "", fieldsUsage)
	flagset.StringVar(&flags.fields, "f", "", fieldsUsage)

	err := flagset.Parse(args)
	if err != nil {
		return nil, nil, err
	}

	return flags, flagset.Args(), nil
}

func main() {
	// default apikey to
	flags, args, err := parseFlags(os.Args[1:])
	if err != nil {
		logger.Fatal(err)
	}

	if len(args) == 0 {
		logger.Fatal("battleaxe need an url")
	}

	url := args[0]

	// let's be tolerent and parse flag after the url
	if len(args) > 1 {
		var afterFlags *appFlags
		afterFlags, args, err = parseFlags(args[1:])
		if err != nil {
			logger.Fatal(err)
		}
		flags = mergeFlags(flags, afterFlags)
	}
	if len(args) > 1 {
		logger.Fatal("battleaxe can parse only one url at a time")
	}

	queryMap := buildQueryMap(flags)

	c := &context{
		url:      url,
		flags:    flags,
		queryMap: queryMap,
	}

	c.action()
}
