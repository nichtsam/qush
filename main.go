package main

import "log"

func main() {
	screen, fini := initScreen()
	defer fini()
	game := NewGame(screen)

	if err := game.Start(); err != nil {
		log.Fatal(err)
	}
}
