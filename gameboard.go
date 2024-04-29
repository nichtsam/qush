package main

import (
	"errors"

	"github.com/gdamore/tcell/v2"
	"golang.org/x/exp/maps"
)

const (
	IDLE buttonState = iota
	WANTED
	PUSHED
)

var (
	ErrButtonNotFound  = errors.New("button not found")
	ErrTriggerNotFound = errors.New("button trigger not found")
)

type (
	Gameboard struct {
		ids     ButtonIds
		buttons Buttons
		layout  Layout
	}
	ButtonTrigger rune
	ButtonIds     map[ButtonTrigger]ButtonId
	ButtonId      int
	Buttons       map[ButtonId]*Button
	Layout        map[ButtonId]ButtonLayout
	ButtonLayout  struct {
		x, y   int
		symbol rune
	}
)

func (gb *Gameboard) HandleKeyEvent(ev *tcell.EventKey) error {
	if ev.Key() == tcell.KeyRune {
		switch r := ev.Rune(); r {
		case ' ':
			gb.reset()
			gb.wantRandomButtons()

		default:
			maybeTrigger := ButtonTrigger(r)
			bid, ok := gb.ids[maybeTrigger]
			if !ok {
				return ErrTriggerNotFound
			}

			err := gb.pushButton(bid)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (gb *Gameboard) draw(s tcell.Screen, force bool) {
	for bid, b := range gb.buttons {
		layout := gb.layout[bid]
		b.draw(s, layout, force)
	}
}

func (gb *Gameboard) reset() {
	for _, b := range gb.buttons {
		b.idle()
	}
}

func (gb *Gameboard) wantRandomButtons() {
	randomIds := generateUniqueRandoms(maps.Keys(gb.buttons))

	for _, id := range randomIds {
		button := gb.buttons[id]
		button.want()
	}
}

func (gb *Gameboard) pushButton(bid ButtonId) error {
	b, ok := gb.buttons[bid]
	if !ok {
		return ErrButtonNotFound
	}
	b.push()
	return nil
}

type (
	Button struct {
		state     buttonState
		prevState buttonState
	}
	buttonState int
)

func (b *Button) draw(s tcell.Screen, layout ButtonLayout, force bool) {
	if b.state == b.prevState && !force {
		return
	}

	x, y, symbol := layout.x, layout.y, layout.symbol

	switch b.state {
	case IDLE:
		s.SetContent(x, y, symbol, nil, tcell.StyleDefault.Foreground(tcell.ColorSeaGreen))
	case WANTED:
		s.SetContent(x, y, symbol, nil, tcell.StyleDefault.Foreground(tcell.ColorDarkOrange))
	case PUSHED:
		s.SetContent(x, y, symbol, nil, tcell.StyleDefault.Foreground(tcell.ColorDarkRed))
	default:
		s.SetContent(x, y, symbol, nil, tcell.StyleDefault.Foreground(tcell.ColorGreenYellow))
	}
}

func (b *Button) idle() {
	b.prevState = b.state
	b.state = IDLE
}

func (b *Button) want() {
	b.prevState = b.state
	b.state = WANTED
}

func (b *Button) push() bool {
	wasWanted := b.state != WANTED
	b.prevState = b.state
	b.state = PUSHED

	return wasWanted
}
