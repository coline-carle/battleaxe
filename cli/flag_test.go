package cli

import "testing"

func compareAppFlags(a *appFlags, b *appFlags) bool {
	if a.locale != b.locale {
		return false
	}
	if a.fields != b.fields {
		return false
	}
	if a.apikey != b.apikey {
		return false
	}
	if a.head != b.head {
		return false
	}
	if a.human != b.human {
		return false
	}
	if a.version != b.version {
		return false
	}
	return true
}

func TestParseCommand(t *testing.T) {
	var tests = []struct {
		in       []string
		outFlags *appFlags
		outURL   string
	}{
		{
			[]string{"url"},
			&appFlags{
				locale:  "",
				fields:  "",
				apikey:  "",
				head:    false,
				human:   false,
				version: false,
			},
			"url",
		},
		{
			[]string{"--apikey=1234", "url"},
			&appFlags{
				locale:  "",
				fields:  "",
				apikey:  "1234",
				head:    false,
				human:   false,
				version: false,
			},
			"url",
		},
		{
			[]string{"url", "--apikey=456"},
			&appFlags{
				locale:  "",
				fields:  "",
				apikey:  "456",
				head:    false,
				human:   false,
				version: false,
			},
			"url",
		},
		{
			[]string{"--apikey=123", "url", "--apikey=456"},
			&appFlags{
				locale:  "",
				fields:  "",
				apikey:  "456",
				head:    false,
				human:   false,
				version: false,
			},
			"url",
		},
	}
	for _, test := range tests {
		outFlags, outURL, err := parseCommand(test.in)
		if err != nil {
			t.Errorf("parseCommand(%q)\n- want:\n%v\n%v\n- got:\n error: %v\n",
				test.in, test.outFlags, test.outURL, err)
		} else if outURL != test.outURL || !compareAppFlags(outFlags, test.outFlags) {
			t.Errorf("parseCommand(%q)\n- want:\nflags=%v\nurl=%v\n- got:\nflags=%v\nurl=%v\n",
				test.in, test.outFlags, test.outURL, outFlags, outURL)
		}
	}
}

func TestParseCommandError(t *testing.T) {
	var tests = [][]string{
		[]string{"--apikey=123"},
		[]string{"--apikey=123", "url1", "url2"},
		[]string{"--apikey=123", "url1", "url2", "--apikey=456"},
	}
	for _, test := range tests {
		outFlags, outURL, err := parseCommand(test)
		if err == nil {
			t.Errorf("parseCommand(%q)\n- want: error- got:\nflags=%v\nurl=%v\n",
				test, outFlags, outURL)
		}
	}
}
