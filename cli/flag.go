package cli

import (
	"errors"
	"flag"
)

type appFlags struct {
	locale string
	fields string
	apikey string

	head    bool
	human   bool
	version bool
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

func parseFlags(args []string) (*appFlags, []string, error) {
	flags := &appFlags{}

	flagset := flag.NewFlagSet("battleaxe", -1)

	localeUsage := "locale"
	flagset.StringVar(&flags.locale, "locale", "", localeUsage)
	flagset.StringVar(&flags.locale, "l", "", localeUsage)

	apikeyUsage := "your personal api key"
	flagset.StringVar(&flags.apikey, "apikey", "", apikeyUsage)
	flagset.StringVar(&flags.apikey, "k", "", apikeyUsage)

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

// parseCommand : parse the command passed to the application
func parseCommand(args []string) (flags *appFlags, url string, err error) {
	flags, args, err = parseFlags(args)
	if err != nil {
		return nil, url, err
	}

	// -version : don't try to parse what's remaining
	if flags.version {
		return flags, "", nil
	}

	if len(args) == 0 {
		return nil, "", errors.New("an url scheme is required")
	}

	url = args[0]
	args = args[1:]

	// let's be tolerent and parse flag after the url
	if len(args) > 0 {
		var afterFlags *appFlags
		afterFlags, args, err = parseFlags(args)
		if err != nil {
			logger.Fatal(err)
		}
		flags = mergeFlags(flags, afterFlags)
	}

	if len(args) > 0 {
		return nil, "", errors.New("one url at a time")
	}

	return flags, url, nil
}
