package main

import "log"

func main() {
	screen, fini := initScreen()
	defer fini()
	game, err := NewGame(screen, "basic")
	if err != nil {
		log.Fatal(err)
	}

	if err = game.Start(); err != nil {
		log.Fatal(err)
	}
}
