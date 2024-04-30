package main

import "github.com/common-nighthawk/go-figure"

type (
	buttonSetCreator func() (ButtonIds, Buttons, Layout)
)

var buttonSetCreators = map[string]buttonSetCreator{
	"raw": func() (ButtonIds, Buttons, Layout) {
		ids := ButtonIds{}
		bs := Buttons{}
		l := Layout{}

		for i := range 26 {
			bid := ButtonId(i)
			upper := rune(65 + i)
			lower := rune(97 + i)
			symbol := upper
			button := &Button{}
			layout := ButtonLayout{
				x: i, y: 0, symbol: [][]rune{{symbol}},
			}
			ids[ButtonTrigger(upper)], ids[ButtonTrigger(lower)] = bid, bid
			bs[bid] = button
			l[bid] = layout
		}

		return ids, bs, l
	},
	"basic": func() (ButtonIds, Buttons, Layout) {
		ids := ButtonIds{}
		bs := Buttons{}

		l := Layout{
			'Q': {0, 0, Runify(figure.NewFigure("Q", "basic", true).Slicify())},
			'W': {10, 0, Runify(figure.NewFigure("W", "basic", true).Slicify())},
			'E': {25, 0, Runify(figure.NewFigure("E", "basic", true).Slicify())},
			'R': {35, 0, Runify(figure.NewFigure("R", "basic", true).Slicify())},
			'T': {45, 0, Runify(figure.NewFigure("T", "basic", true).Slicify())},
			'Y': {55, 0, Runify(figure.NewFigure("Y", "basic", true).Slicify())},
			'U': {65, 0, Runify(figure.NewFigure("U", "basic", true).Slicify())},
			'I': {75, 0, Runify(figure.NewFigure("I", "basic", true).Slicify())},
			'O': {85, 0, Runify(figure.NewFigure("O", "basic", true).Slicify())},
			'P': {95, 0, Runify(figure.NewFigure("P", "basic", true).Slicify())},

			'A': {5, 7, Runify(figure.NewFigure("A", "basic", true).Slicify())},
			'S': {15, 7, Runify(figure.NewFigure("S", "basic", true).Slicify())},
			'D': {25, 7, Runify(figure.NewFigure("D", "basic", true).Slicify())},
			'F': {35, 7, Runify(figure.NewFigure("F", "basic", true).Slicify())},
			'G': {45, 7, Runify(figure.NewFigure("G", "basic", true).Slicify())},
			'H': {55, 7, Runify(figure.NewFigure("H", "basic", true).Slicify())},
			'J': {65, 7, Runify(figure.NewFigure("J", "basic", true).Slicify())},
			'K': {75, 7, Runify(figure.NewFigure("K", "basic", true).Slicify())},
			'L': {85, 7, Runify(figure.NewFigure("L", "basic", true).Slicify())},

			'Z': {10, 14, Runify(figure.NewFigure("Z", "basic", true).Slicify())},
			'X': {20, 14, Runify(figure.NewFigure("X", "basic", true).Slicify())},
			'C': {30, 14, Runify(figure.NewFigure("C", "basic", true).Slicify())},
			'V': {40, 14, Runify(figure.NewFigure("V", "basic", true).Slicify())},
			'B': {50, 14, Runify(figure.NewFigure("B", "basic", true).Slicify())},
			'N': {60, 14, Runify(figure.NewFigure("N", "basic", true).Slicify())},
			'M': {71, 14, Runify(figure.NewFigure("M", "basic", true).Slicify())},
		}

		for i := range 26 {
			upper := rune(65 + i)
			lower := rune(97 + i)
			bid := ButtonId(upper)
			button := &Button{}
			ids[ButtonTrigger(upper)], ids[ButtonTrigger(lower)] = bid, bid
			bs[bid] = button
		}

		return ids, bs, l
	},
}
