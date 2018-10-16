
[![CircleCI](https://circleci.com/gh/coline-carle/battleaxe.svg?style=svg)](https://circleci.com/gh/coline-carle/battleaxe)

opiniated terminal client for blizzard.com community API

## configuration
```sh
export BLIZZARD_CLIENT_ID="..."
export BLIZZARD_CLIENT_SECRET="..."
```

## The storyline

```
battleaxe  us://wow/achievement/2144
battleaxe wow://achievement/2144
wowaxe  us://achievement/2144
wowaxe achievment://2144
```

## binaries
- battleaxe (All games endpoints)
- wowaxe (World of Warcraft enpoint)
- scaxe (Starcraft 2 endpoint)
- daxe (Diablo 3 endpoint)
## Usage

```
USAGE:
  battleaxe [OPTIONS] scheme://path/to/resource [MORE OPTIONS]

DESCRIPTION:
        examples of urls:

        Golden rule: Flags that modify query options, always take precedence
        over the query option of the url. In case of multiple definition of the same
        flag: The rightest flag win.

QUERY OPTIONS:
  client, K     blizzard API client id, environnement variable BLIZZARD_CLIENT_ID is used by default if unset
  secret, S     blizzard API client secret, environnement variable BLIZZARD_CLIENT_SECRET is used by default if unset
  fields, F     set optional fields for the requested endpoint
  locale, L     game locale example: en_US

GLOBAL OPTIONS:
  head, I       print headers instead of body
  human, C      humanize response with color and indentation
  pretty        humanize response with color and indentation (same as human)
  version, V    show version
  dry, D        print the url that would be request instead of fetching it
  help, usage   show this help

SEE ALSO:
        battleaxe, wowaxe, scaxe, daxe

VERSION:
  0.0.10

```

## Releases

https://github.com/coline-carle/battleaxe/releases

or

```
go get -u github.com/coline-carle/battleaxe/cmd/battleaxe
go get -u github.com/coline-carle/battleaxe/cmd/wowaxe
go get -u github.com/coline-carle/battleaxe/cmd/daxe
go get -u github.com/coline-carle/battleaxe/cmd/scaxe
```


## License

MIT License

Copyright (c) 2018 Coline Carle

