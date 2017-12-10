package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/wow-sweetlie/battleaxe/battle"

	"github.com/fatih/color"
	"github.com/nwidger/jsoncolor"
)

type appFlags struct {
	locale string
	fields string
	apikey string

	head    bool
	human   bool
	version bool
}

type context struct {
	queryMap map[string]string
	flags    *appFlags
	url      string
}

const version = "0.0.1"

var (
	colorField = color.New(color.FgBlue, color.Bold).SprintFunc()
	colorValue = color.New().SprintFunc()
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stderr, "", 0)
}

func firstString(a string, b string) string {
	if a == "" {
		return b
	}
	return a
}

func mergeFlags(flags *appFlags, afterFlags *appFlags) *appFlags {
	return &appFlags{
		locale: firstString(afterFlags.locale, flags.locale),
		fields: firstString(afterFlags.fields, flags.fields),
		apikey: firstString(afterFlags.apikey, flags.apikey),
		head:   flags.head || afterFlags.head,
		human:  flags.human || afterFlags.human,
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
	if c.flags.human {
		b := new(bytes.Buffer)
		b.ReadFrom(resp.Body)
		f := jsoncolor.NewFormatter()
		err := f.Format(os.Stdout, b.Bytes())
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}
	_, err := io.Copy(os.Stdout, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func (c *context) displayHeader(resp *http.Response) {
	for field, value := range resp.Header {
		formatedValue := strings.Join(value, ",")
		fmt.Printf("%s: %s\n", colorField(field), colorValue(formatedValue))
	}
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

	if c.flags.head {
		c.displayHeader(resp)
		os.Exit(0)
	}

	err = c.display(resp)
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

	humanUsage := "humanize the output with some color and proper indenting"
	flagset.BoolVar(&flags.human, "color", false, humanUsage)
	flagset.BoolVar(&flags.human, "c", false, humanUsage)
	flagset.BoolVar(&flags.human, "h", false, humanUsage)

	headUsage := "show headers"
	flagset.BoolVar(&flags.head, "head", false, headUsage)
	flagset.BoolVar(&flags.head, "I", false, headUsage)

	versionUsage := version
	flagset.BoolVar(&flags.version, "version", false, versionUsage)
	flagset.BoolVar(&flags.version, "V", false, versionUsage)

	err := flagset.Parse(args)
	if err != nil {
		return nil, nil, err
	}

	return flags, flagset.Args(), nil
}

func showVersion() {
	fmt.Printf("v%s\n", version)
	os.Exit(0)
}

func main() {
	// default apikey to
	flags, args, err := parseFlags(os.Args[1:])
	if err != nil {
		logger.Fatal(err)
	}

	if flags.version {
		showVersion()
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
