package cli

import (
	"errors"
	"flag"
)

type appFlags struct {
	locale       string
	fields       string
	clientID     string
	clientSecret string

	head    bool
	human   bool
	version bool
	dry     bool
	help    bool
}

func firstString(a string, b string) string {
	if a == "" {
		return b
	}
	return a
}

func mergeFlags(flags *appFlags, afterFlags *appFlags) *appFlags {
	return &appFlags{
		locale:       firstString(afterFlags.locale, flags.locale),
		fields:       firstString(afterFlags.fields, flags.fields),
		clientID:     firstString(afterFlags.clientID, flags.clientID),
		clientSecret: firstString(afterFlags.clientSecret, flags.clientSecret),
		head:         flags.head || afterFlags.head,
		human:        flags.human || afterFlags.human,
		version:      flags.version || afterFlags.version,
		dry:          flags.dry || afterFlags.dry,
		help:         flags.help || afterFlags.help,
	}
}

func buildFlagset(flags *appFlags) *flag.FlagSet {
	flagset := flag.NewFlagSet("battleaxe", flag.ExitOnError)

	flagset.StringVar(&flags.locale, "locale", "", "")
	flagset.StringVar(&flags.locale, "L", "", "")

	flagset.StringVar(&flags.clientID, "client", "", "")
	flagset.StringVar(&flags.clientID, "K", "", "")

	flagset.StringVar(&flags.clientSecret, "secret", "", "")
	flagset.StringVar(&flags.clientSecret, "S", "", "")

	flagset.StringVar(&flags.fields, "fields", "", "")
	flagset.StringVar(&flags.fields, "F", "", "")

	flagset.BoolVar(&flags.human, "pretty", false, "")
	flagset.BoolVar(&flags.human, "human", false, "")
	flagset.BoolVar(&flags.human, "C", false, "")

	flagset.BoolVar(&flags.head, "head", false, "")
	flagset.BoolVar(&flags.head, "I", false, "")

	flagset.BoolVar(&flags.version, "version", false, "")
	flagset.BoolVar(&flags.version, "V", false, "")

	flagset.BoolVar(&flags.dry, "dry", false, usageDry)
	flagset.BoolVar(&flags.dry, "D", false, usageDry+" (shorthand)")

	flagset.BoolVar(&flags.help, "help", false, usageHelp)
	flagset.BoolVar(&flags.help, "usage", false, usageHelp)

	flagset.Usage = func() {
		PrintHelp()
	}

	return flagset
}

func parseFlags(args []string) (*appFlags, []string, error) {
	flags := &appFlags{}

	flagset := buildFlagset(flags)

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

	// don't try to parse what remains
	if flags.version || flags.help {
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
			return nil, "", err
		}
		flags = mergeFlags(flags, afterFlags)
	}

	if len(args) > 0 {
		return nil, "", errors.New("one url at a time")
	}

	return flags, url, nil
}
