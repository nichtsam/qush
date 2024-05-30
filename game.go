package main

import (
	"errors"

	"github.com/gdamore/tcell/v2"
)

type Game struct {
	screen    tcell.Screen
	audio     Audio
	gameboard *Gameboard
}

var ErrInvalidButtonSetName = errors.New("invalid button set name")

func NewGame(screen tcell.Screen, buttonSetName string) (*Game, error) {
	audio := newAudio()
	gameboard := &Gameboard{}

	buttonSetCreator, ok := buttonSetCreators[buttonSetName]
	if !ok {
		return nil, ErrInvalidButtonSetName
	}
	gameboard.ids, gameboard.buttons, gameboard.layout = buttonSetCreator()

	return &Game{
		screen,
		audio,
		gameboard,
	}, nil
}

func (g *Game) Start() error {
	if err := g.audio.register("pop", "./assets/sounds/pop.mp3"); err != nil {
		return err
	}

	g.draw(true)

	for {
		ev := g.screen.PollEvent()
		gb := g.gameboard

		switch ev := ev.(type) {
		case *tcell.EventResize:
			g.screen.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				return nil
			}

			if ev.Key() == tcell.KeyRune {
				switch r := ev.Rune(); r {
				case ' ':
					gb.NewRound()

				default:
					_ = g.audio.play("pop")
					maybeTrigger := ButtonTrigger(r)
					bid, ok := gb.ids[maybeTrigger]
					if !ok {
						return ErrTriggerNotFound
					}

					err := gb.PushButton(bid)
					if err != nil {
						return err
					}
				}
			}

		}

		g.draw(false)
	}
}

func (g *Game) draw(force bool) {
	g.gameboard.draw(g.screen, force)

	g.screen.Show()
}
