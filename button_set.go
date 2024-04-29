package main

type (
	buttonsCreator func() (ButtonIds, Buttons, Layout)
)

var buttonSets = map[string]buttonsCreator{
	"default": func() (ButtonIds, Buttons, Layout) {
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
				x: i, y: 0, symbol: symbol,
			}
			ids[ButtonTrigger(upper)], ids[ButtonTrigger(lower)] = bid, bid
			bs[bid] = button
			l[bid] = layout
		}

		return ids, bs, l
	},
}
