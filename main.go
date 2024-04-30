package main

import (
	"flag"
	"log"
)

func main() {
	buttonSetName := flag.String(
		"buttonset",
		"basic",
		"the button set the gameboard is going to use,\nbutton set to use: basic, raw.",
	)
	flag.Parse()

	screen, fini := initScreen()
	defer fini()
	game, err := NewGame(screen, *buttonSetName)
	if err != nil {
		fini()
		log.Fatal(err)
	}

	if err = game.Start(); err != nil {
		fini()
		log.Fatal(err)
	}
}
