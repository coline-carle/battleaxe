# battleaxe
[![CircleCI](https://circleci.com/gh/coline-carle/battleaxe.svg?style=svg)](https://circleci.com/gh/coline-carle/battleaxe)

battleaxe is an opiniated terminal client for blizzard.com community API

## configuration
export BLIZZARD_CLIENT_ID="..."
export BLIZZARD_CLIENT_SECRET="..."


## The storyline

```
battleaxe  us://wow/achievement/2144
````

```
battleaxe wow://achievement/2144
````

```
wowaxe  us://achievement/2144
````

## binaries

battleaxe come with few executables for querying battlenet API

Here are their kind names:

- battleaxe (All games endpoints)
- wowaxe (World of Warcraft enpoint)
- scaxe (Starcraft 2 endpoint)
- daxe (Diablo 3 endpoint)

Be gentle with them, I crafted them with love.

## Usage

```
NAME:
  battleaxe

USAGE:
  battleaxe [OPTIONS] scheme://path/to/resource [MORE OPTIONS]

DESCRIPTION:
        examples of urls:
          https://us.api.blizzard.com/wow/achievement/2144?locale=en_US&apikey=APIKEY
          us://wow/achievement/2144
          wow://achievement/2144

        Golden rule: Flags that modify query options, always take precedence
        over the query option of the url. In case of multiple definition of the same
        flag: The rightest flag win.

QUERY OPTIONS:
  apikey, k     your personal api key variable env BATTLENET_CLIENT_ID is used by default if set
  fields, f     set fields for endpoint that accept ones
  locale, l     game locale ex: en_US

GLOBAL OPTIONS:
  head, I       print headers instead of body
  color, C      humanize response with color and indentation
  version, V    show version
  dry, D        only print the url that would be requested
  help, usage   show this help

SEE ALSO:
        battleaxe, wowaxe, scaxe, daxe

VERSION:
  0.0.1
```
## Releases

https://github.com/coline-carle/battleaxe/releases

if go is installed and configured, you can get master of one one the tool like
this :

* go get -U github.com/coline-carle/battleaxe/cmd/battleaxe
* go get -U github.com/coline-carle/battleaxe/cmd/wowaxe
* go get -U github.com/coline-carle/battleaxe/cmd/daxe
* go get -U github.com/coline-carle/battleaxe/cmd/scaxe



## License

MIT License

Copyright (c) 2017 Sweetlie <Colin Carle>

