package battle

import "testing"

func TestParseURLWithGame(t *testing.T) {
	var tests = []struct {
		in   string
		out  string
		game Game
	}{
		{
			"eu://achievement/2144?locale=en_US&apikey=APIKEY",
			"https://eu.api.blizzard.com/wow/achievement/2144?apikey=APIKEY&locale=en_US",
			WoW,
		},
		{
			"achievement://2144?locale=en_US&apikey=APIKEY",
			"https://us.api.blizzard.com/wow/achievement/2144?apikey=APIKEY&locale=en_US",
			WoW,
		},
		{
			"character://Archimonde/Sweetlie",
			"https://us.api.blizzard.com/wow/character/Archimonde/Sweetlie",
			WoW,
		},
	}
	for _, test := range tests {
		out, err := ParseURL(test.in, nil, test.game)
		if err != nil {
			t.Errorf("parseURL(%q, nil)\n- want:\n %v\n- got:\n error: %v\n", test.in, test.out, err)
		} else if out != test.out {
			t.Errorf("parseURL(%q, nil)\n- want:\n %v\n- got:\n %v\n", test.in, test.out, out)
		}
	}
}

func TestParseURL(t *testing.T) {
	var tests = []struct {
		in  string
		out string
	}{
		{
			"https://us.api.blizzard.com/wow/achievement/2144?locale=en_US&apikey=APIKEY",
			"https://us.api.blizzard.com/wow/achievement/2144?apikey=APIKEY&locale=en_US",
		},
		{
			"http://us.api.blizzard.com/wow/achievement/2144",
			"https://us.api.blizzard.com/wow/achievement/2144",
		},
		{
			"us://wow/achievement/2144?locale=en_US&apikey=APIKEY",
			"https://us.api.blizzard.com/wow/achievement/2144?apikey=APIKEY&locale=en_US",
		},
		{
			"us://wow/achievement/2144",
			"https://us.api.blizzard.com/wow/achievement/2144",
		},
		{
			"eu://wow/achievement/2144",
			"https://eu.api.blizzard.com/wow/achievement/2144",
		},
		{
			"wow://achievement/2144",
			"https://us.api.blizzard.com/wow/achievement/2144",
		},
	}
	for _, test := range tests {
		out, err := ParseURL(test.in, nil, Any)
		if err != nil {
			t.Errorf("parseURL(%q, nil)\n- want:\n %v\n- got:\n error: %v\n", test.in, test.out, err)
		} else if out != test.out {
			t.Errorf("parseURL(%q, nil)\n- want:\n %v\n- got:\n %v\n", test.in, test.out, out)
		}
	}
}

func TestParseURLWithOverride(t *testing.T) {
	var tests = []struct {
		in       string
		out      string
		override map[string]string
	}{
		{
			"https://us.api.blizzard.com/wow/achievement/2144?locale=en_US&apikey=APIKEY",
			"https://us.api.blizzard.com/wow/achievement/2144?apikey=APIKEY&locale=fr_FR",
			map[string]string{
				"locale": "fr_FR",
			},
		},
		{
			"https://us.api.blizzard.com/wow/achievement/2144?apikey=APIKEY",
			"https://us.api.blizzard.com/wow/achievement/2144?apikey=APIKEY&locale=fr_FR",
			map[string]string{
				"locale": "fr_FR",
			},
		},
		{
			"https://us.api.blizzard.com/wow/achievement/2144?apikey=APIKEY",
			"https://us.api.blizzard.com/wow/achievement/2144?apikey=mykey&locale=fr_FR",
			map[string]string{
				"locale": "fr_FR",
				"apikey": "mykey",
			},
		},
	}

	for _, test := range tests {
		out, err := ParseURL(test.in, test.override, Any)
		if err != nil {
			t.Errorf("parseURL(%q, %q)\n- want:\n %v\n- got error:\n %v\n", test.in, test.override, test.out, err)
		} else if out != test.out {
			t.Errorf("parseURL(%q, %q)\n- want:\n %v\n- got:\n %v\n", test.in, test.override, test.out, out)
		}
	}
}
