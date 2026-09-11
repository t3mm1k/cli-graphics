package widgets

import (
	"cli-graphics/engine"
	"unicode/utf8"
)

type Button struct {
	X, Y int
	W, H int

	text   string
	Buffer *engine.Buffer

	isFocused bool
	onClick   func()
}

func NewButton(x, y int, text string, onClick func()) *Button {
	width := utf8.RuneCountInString(text)
	width, height := GetActuallySize(true, width, 1)

	buf := engine.NewBuffer(width, height)

	return &Button{
		text:      text,
		X:         x,
		Y:         y,
		W:         width,
		H:         height,
		Buffer:    buf,
		isFocused: false,
		onClick:   onClick,
	}
}

func (b *Button) Coords() (x, y int) {
	return b.X, b.Y
}

func (b *Button) GetBuffer() [][]rune {
	return b.Buffer.GetObjects()
}

func (b *Button) SetFocus(focused bool) {
	b.isFocused = focused
}

func (b *Button) IsFocused() bool {
	return b.isFocused
}

func (b *Button) HandleKey(key string) bool {
	if key == "Enter" || key == " " {
		if b.onClick != nil {
			b.onClick()
		}
		return true
	}
	return false
}

func (b *Button) Render() {
	var hLine, vLine rune
	var tl, tr, bl, br rune

	if b.isFocused {
		hLine, vLine = '═', '║'
		tl, tr, bl, br = '╔', '╗', '╚', '╝'
	} else {
		hLine, vLine = '─', '│'
		tl, tr, bl, br = '┌', '┐', '└', '┘'
	}

	for x := 1; x < b.W-1; x++ {
		b.Buffer.Data[0][x] = hLine
		b.Buffer.Data[b.H-1][x] = hLine
	}

	for y := 1; y < b.H-1; y++ {
		b.Buffer.Data[y][0] = vLine
		b.Buffer.Data[y][b.W-1] = vLine
	}

	b.Buffer.Data[0][0] = tl
	b.Buffer.Data[0][b.W-1] = tr
	b.Buffer.Data[b.H-1][0] = bl
	b.Buffer.Data[b.H-1][b.W-1] = br

	textW := b.W - 2
	textH := 1

	for i, let := range b.text {
		if i >= textW {
			break
		}
		b.Buffer.Data[textH][i+1] = let
	}
}
