package main

import (
	"os"

	"github.com/wow-sweetlie/battleaxe/battle"
	"github.com/wow-sweetlie/battleaxe/cli"
)

func main() {
	cli.Run(battle.SC2, os.Args)
}
