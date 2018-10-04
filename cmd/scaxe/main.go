package main

import (
	"os"

	"github.com/coline-carle/battleaxe/battle"
	"github.com/coline-carle/battleaxe/cli"
)

func main() {
	cli.Run(battle.SC2, os.Args)
}
