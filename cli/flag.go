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
		locale:  firstString(afterFlags.locale, flags.locale),
		fields:  firstString(afterFlags.fields, flags.fields),
		apikey:  firstString(afterFlags.apikey, flags.apikey),
		head:    flags.head || afterFlags.head,
		human:   flags.human || afterFlags.human,
		version: flags.version || afterFlags.version,
		dry:     flags.dry || afterFlags.dry,
		help:    flags.help || afterFlags.help,
	}
}

func parseFlags(args []string) (*appFlags, []string, error) {
	flags := &appFlags{}

	flagset := flag.NewFlagSet("battleaxe", -1)

	// I do not use usage string here see : help.go
	usage := ""
	flagset.StringVar(&flags.locale, "locale", "", usage)
	flagset.StringVar(&flags.locale, "l", "", usage)

	flagset.StringVar(&flags.apikey, "apikey", "", usage)
	flagset.StringVar(&flags.apikey, "k", "", usage)

	flagset.StringVar(&flags.fields, "fields", "", usage)
	flagset.StringVar(&flags.fields, "f", "", usage)

	flagset.BoolVar(&flags.human, "color", false, usage)
	flagset.BoolVar(&flags.human, "human", false, usage)
	flagset.BoolVar(&flags.human, "C", false, usage)

	flagset.BoolVar(&flags.head, "head", false, usage)
	flagset.BoolVar(&flags.head, "I", false, usage)

	flagset.BoolVar(&flags.version, "version", false, usage)
	flagset.BoolVar(&flags.version, "V", false, usage)

	flagset.BoolVar(&flags.dry, "dry", false, usage)
	flagset.BoolVar(&flags.dry, "D", false, usage)

	flagset.BoolVar(&flags.help, "help", false, usage)
	flagset.BoolVar(&flags.help, "usage", false, usage)

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
			return nil, "", err
		}
		flags = mergeFlags(flags, afterFlags)
	}

	if len(args) > 0 {
		return nil, "", errors.New("one url at a time")
	}

	return flags, url, nil
}
