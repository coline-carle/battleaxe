package battle

import "testing"

func TestParseURL(t *testing.T) {
	var tests = []struct {
		in  string
		out string
	}{
		{
			"https://us.api.battle.net/wow/achievement/2144?locale=en_US&apikey=<APIKEY>",
			"https://us.api.battle.net/wow/achievement/2144?locale=en_US",
		},
		{
			"http://us.api.battle.net/wow/achievement/2144",
			"https://us.api.battle.net/wow/achievement/2144",
		},
		{
			"us://wow/achievement/2144?locale=en_US&apikey=<APIKEY>",
			"https://us.api.battle.net/wow/achievement/2144?locale=en_US",
		},
		{
			"us://wow/achievement/2144",
			"https://us.api.battle.net/wow/achievement/2144",
		},
		{
			"eu://wow/achievement/2144",
			"https://eu.api.battle.net/wow/achievement/2144",
		},
		{
			"wow://achievement/2144",
			"https://us.api.battle.net/wow/achievement/2144",
		},
	}
	for _, test := range tests {
		out, err := ParseURL(test.in)
		if err != nil {
			t.Errorf("parseURL(%q) want: %v got error: %v", test.in, test.out, err)
		} else if out != test.out {
			t.Errorf("parseURL(%q)\n- want:\n %v\n- got:\n %v\n", test.in, test.out, out)
		}
	}
}
