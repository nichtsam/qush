package main

import (
	"errors"

	"github.com/gdamore/tcell/v2"
)

type Game struct {
	screen    tcell.Screen
	gameboard *Gameboard
}

var ErrInvalidButtonSetName = errors.New("invalid button set name")

func NewGame(screen tcell.Screen, buttonSetName string) (*Game, error) {
	gameboard := &Gameboard{}

	buttonSetCreator, ok := buttonSetCreators[buttonSetName]
	if !ok {
		return nil, ErrInvalidButtonSetName
	}
	gameboard.ids, gameboard.buttons, gameboard.layout = buttonSetCreator()

	return &Game{
		screen,
		gameboard,
	}, nil
}

func (g *Game) Start() error {
	g.draw(true)

	for {
		ev := g.screen.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventResize:
			g.screen.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				return nil
			}

			_ = g.gameboard.HandleKeyEvent(ev)

		}

		g.draw(false)
	}
}

func (g *Game) draw(force bool) {
	g.gameboard.draw(g.screen, force)

	g.screen.Show()
}
