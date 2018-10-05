package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCommand(t *testing.T) {
	var tests = []struct {
		in            []string
		expectedFlags *appFlags
		expectedURL   string
	}{
		{
			[]string{"url"},
			&appFlags{
				locale:       "",
				fields:       "",
				clientID:     "",
				clientSecret: "",
				head:         false,
				human:        false,
				version:      false,
			},
			"url",
		},
		{
			[]string{"--client=1234", "url"},
			&appFlags{
				locale:       "",
				fields:       "",
				clientID:     "1234",
				clientSecret: "",
				head:         false,
				human:        false,
				version:      false,
			},
			"url",
		},
		{
			[]string{"url", "--client=456", "--secret=secret"},
			&appFlags{
				locale:       "",
				fields:       "",
				clientID:     "456",
				clientSecret: "secret",
				head:         false,
				human:        false,
				version:      false,
			},
			"url",
		},
		{
			[]string{"--client=123", "--secret=secret", "url", "--client=456"},
			&appFlags{
				locale:       "",
				fields:       "",
				clientID:     "456",
				clientSecret: "secret",
				head:         false,
				human:        false,
				version:      false,
			},
			"url",
		},
	}
	for _, test := range tests {
		flags, url, err := parseCommand(test.in)
		if err != nil {
			t.Errorf("parseCommand(%q)\n- want:\n%v\n%v\n- got:\n error: %v\n",
				test.in, test.expectedFlags, test.expectedURL, err)
		}
		assert.Equal(t, test.expectedURL, url)
		assert.Equal(t, test.expectedFlags, flags)
	}
}

func TestParseCommandError(t *testing.T) {
	var tests = [][]string{
		{"--client=123"},
		{"--client=123", "url1", "url2"},
		{"--client=123", "url1", "url2", "--client=456"},
	}
	for _, test := range tests {
		outFlags, outURL, err := parseCommand(test)
		if err == nil {
			t.Errorf("parseCommand(%q)\n- want: error- got:\nflags=%v\nurl=%v\n",
				test, outFlags, outURL)
		}
	}
}
