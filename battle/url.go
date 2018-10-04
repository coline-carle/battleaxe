package battle

import (
	"errors"
	"net/url"
	"strings"
)

const (
	// USHost : United-States endpoint
	usHost = "us.api.blizzard.com"

	// EUHost : Europe endpoint
	euHost = "eu.api.blizzard.com"

	// KRHost : Korea endpoint
	krHost = "kr.api.blizzard.com"

	// TWHost : Taiwan endpoint
	twHost = "tw.api.blizzard.com"

	// CNHost : China endpoint
	cnHost = "api.battlenet.com.cn"

	// SEAHost : South East Asia endpoint
	seaHost = "sea.api.blizzard.com"
)

// knwon and parsed Fields
const (
	queryLocale = "locale"
	queryFields = "fields"
	queryAPIKey = "apikey"
)

const (
	// DefaultScheme : scheme use by default in requests
	DefaultScheme = "https"
	// DefaultHost : if no parameter change the endpoint US is default
	// endpoint
	DefaultHost = usHost
)

const (
	// WowPath : World of Wacraft path root
	wowPath = "/wow"
	// D3 : Diablo 3 root path
	d3Path = "/d3"
	// SC2 : Starcraft 2 root path
	sc2Path = "/sc2"
)

type parser struct {
	inURL  *url.URL
	outURL *url.URL
}

// ParseURL parse an input url
func ParseURL(clientURL string, queryOverride map[string]string, game Game) (string, error) {
	u, err := url.Parse(clientURL)
	if err != nil {
		return "", err
	}

	outURL := &url.URL{
		Scheme:   DefaultScheme,
		Host:     u.Host,
		Path:     u.Path,
		RawQuery: u.RawQuery,
	}

	p := &parser{
		inURL:  u,
		outURL: outURL,
	}

	switch game {
	case WoW:
		err = p.parseGameScheme(wowPath)
	case D3:
		err = p.parseGameScheme(d3Path)
	case SC2:
		err = p.parseGameScheme(sc2Path)
	default:
		err = p.parseGeneralScheme()
	}

	if err != nil {
		return "", err
	}
	p.parseQuery(queryOverride)
	return p.outURL.String(), nil
}

func concatPath(gamePath string, path string) string {
	s := []string{gamePath, path}
	return strings.Join(s, "")
}

func transformHostToPath(host string, path string) string {
	s := []string{"/", host, path}
	return strings.Join(s, "")
}

func (p *parser) parseQuery(queryOverride map[string]string) {
	v := p.inURL.Query()
	for key, value := range queryOverride {
		v.Set(key, value)
	}
	p.outURL.RawQuery = v.Encode()
}

func (p *parser) parseGameScheme(gamePath string) error {
	scheme := strings.ToLower(p.inURL.Scheme)
	switch scheme {
	case "eu":
		p.outURL.Host = euHost
		p.outURL.Path = strings.Join([]string{gamePath, "/", p.inURL.Host, p.inURL.Path}, "")
	case "us":
		p.outURL.Host = usHost
		p.outURL.Path = strings.Join([]string{gamePath, "/", p.inURL.Host, p.inURL.Path}, "")
	case "usa":
		p.outURL.Host = usHost
		p.outURL.Path = strings.Join([]string{gamePath, "/", p.inURL.Host, p.inURL.Path}, "")
	case "kr":
		p.outURL.Host = krHost
		p.outURL.Path = strings.Join([]string{gamePath, "/", p.inURL.Host, p.inURL.Path}, "")
	case "tw":
		p.outURL.Host = twHost
		p.outURL.Path = strings.Join([]string{gamePath, "/", p.inURL.Host, p.inURL.Path}, "")
	case "sea":
		p.outURL.Host = seaHost
		p.outURL.Path = strings.Join([]string{gamePath, "/", p.inURL.Host, p.inURL.Path}, "")

	case "https":
		p.outURL.Host = p.inURL.Host
	case "http":
		p.outURL.Host = p.inURL.Host
	default:
		path := strings.Join([]string{gamePath, scheme, p.inURL.Host}, "/")
		p.outURL.Path = concatPath(path, p.inURL.Path)
		p.outURL.Host = DefaultHost
	}
	return nil
}

// setup the game and entrypoint part of the url, based on input shceme
func (p *parser) parseGeneralScheme() error {
	scheme := strings.ToLower(p.inURL.Scheme)

	switch scheme {

	// url is game://path type (default endpoint to US)
	case "wow":
		path := transformHostToPath(p.inURL.Host, p.inURL.Path)
		p.outURL.Path = concatPath(wowPath, path)
		p.outURL.Host = DefaultHost
	case "sc2":
		path := transformHostToPath(p.inURL.Host, p.inURL.Path)
		p.outURL.Path = concatPath(sc2Path, path)
		p.outURL.Host = DefaultHost
	case "d3":
		path := transformHostToPath(p.inURL.Host, p.inURL.Path)
		p.outURL.Path = concatPath(sc2Path, path)
		p.outURL.Host = DefaultHost

		// url is region://game/path type
	case "eu":
		p.outURL.Host = euHost
		p.outURL.Path = concatPath(p.inURL.Host, p.inURL.Path)
	case "us":
		p.outURL.Host = usHost
		p.outURL.Path = concatPath(p.inURL.Host, p.inURL.Path)
	case "usa":
		p.outURL.Host = usHost
		p.outURL.Path = concatPath(p.inURL.Host, p.inURL.Path)
	case "kr":
		p.outURL.Host = krHost
		p.outURL.Path = concatPath(p.inURL.Host, p.inURL.Path)
	case "tw":
		p.outURL.Host = twHost
		p.outURL.Path = concatPath(p.inURL.Host, p.inURL.Path)
	case "sea":
		p.outURL.Host = seaHost
		p.outURL.Path = concatPath(p.inURL.Host, p.inURL.Path)

	case "https":
		p.outURL.Host = p.inURL.Host
	case "http":
		p.outURL.Host = p.inURL.Host
	default:
		return errors.New("invalid scheme")
	}

	return nil
}
