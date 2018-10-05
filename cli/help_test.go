package cli

import (
	"flag"
	"testing"
)

func isFlagHelpDefinedIn(flags []FlagHelp, flagName string) bool {
	for _, aflag := range flags {
		for _, name := range aflag.Names {
			if flagName == name {
				return true
			}
		}
	}
	return false
}

func isFlagHelpDefined(flagName string) bool {
	return isFlagHelpDefinedIn(generalFlags, flagName) || isFlagHelpDefinedIn(queryFlags, flagName)
}

func TestHelp(t *testing.T) {
	visitor := func(f *flag.Flag) {
		if !isFlagHelpDefined(f.Name) {
			t.Errorf("help is not defined for flag: %s\n", f.Name)
		}
	}
	appFlags := &appFlags{}

	flagset := buildFlagset(appFlags)
	flagset.VisitAll(visitor)
}
