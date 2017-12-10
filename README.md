# battleaxe
[![circleCI](https://circleci.com/gh/wow-sweetlie/battleaxe.svg?style=svg)](https://circleci.com/gh/wow-sweetlie/battleaxe)

battleaxe is an opiniated terminal client for Battle.NET community API

## The storyline

* it can ask raw url, replacing APIKEY placeholder by your pre-configured
  BATTLENET_CLIENT_ID env variable

```
battleaxe  "https://us.api.battle.net/wow/achievement/2144?locale=en_US&apikey=APIKEY"
```

* Maybe you haven't set-up yet the env variale, so let's add an argument

```
battleaxe  "https://us.api.battle.net/wow/achievement/2144?locale=en_US&apikey=APIKEY" --apikey myapikey
```

* Nah let's put apikey flag first

```
battleaxe -k myapikey  "https://us.api.battle.net/wow/achievement/2144?locale=en_US&apikey=APIKEY"
```

* To verbose let's use the axe

```
battleaxe -k myapikey  us://wow/achievement/2144
````

* Shorter ! Can we do better ? Region doesn't matter to get the English
  locale of an achievement, let's use the default endpoint (US)

```
battleaxe -k myapikey  wow://achievement/2144
````

* Come to think of it, I'm only using World of Warcraft API

```
wowaxe -k myapikey  us://achievement/2144
````

* It's time to setup the env var

```
export BATTLENET_CLIENT_ID=myapikey
````

* End of Story

```
wowaxe achievment://2144
```


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
          https://us.api.battle.net/wow/achievement/2144?locale=en_US&apikey=APIKEY
          us://wow/achievement/2144
          wow://achievement/2144

        Golden rule: Flags that modify query options, always take precedance
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

## License

MIT License

Copyright (c) 2017 Sweetlie <Colin Carle>

