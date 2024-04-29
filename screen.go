package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
)

func initScreen() (tcell.Screen, func()) {
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatal(err)
	}
	if err = s.Init(); err != nil {
		log.Fatal(err)
	}

	return s, func() {
		s.Fini()
		maybePanic := recover()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
}
