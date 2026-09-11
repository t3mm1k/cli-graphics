package widgets

import (
	"cli-graphics/engine"
	"unicode/utf8"
)

type Button struct {
	x, y int
	w, h int

	text   string
	Buffer *engine.Buffer

	isFocused bool
	OnClick   func()
}

func (b *Button) SetText(newText string) {
	b.text = newText

	width := utf8.RuneCountInString(newText)
	b.w, b.h = GetActuallySize(true, width, 1)
	b.Buffer = engine.NewBuffer(b.w, b.h)
}

func NewButton(x, y int, text string, onClick func()) *Button {
	width := utf8.RuneCountInString(text)
	width, height := GetActuallySize(true, width, 1)

	buf := engine.NewBuffer(width, height)

	return &Button{
		text:      text,
		x:         x,
		y:         y,
		w:         width,
		h:         height,
		Buffer:    buf,
		isFocused: false,
		OnClick:   onClick,
	}
}

func (b *Button) GetCoords() (x, y int) {
	return b.x, b.y
}

func (b *Button) GetSize() (w, h int) {
	return b.w, b.h
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
		if b.OnClick != nil {
			b.OnClick()
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

	for x := 1; x < b.w-1; x++ {
		b.Buffer.Data[0][x] = hLine
		b.Buffer.Data[b.h-1][x] = hLine
	}

	for y := 1; y < b.h-1; y++ {
		b.Buffer.Data[y][0] = vLine
		b.Buffer.Data[y][b.w-1] = vLine
	}

	b.Buffer.Data[0][0] = tl
	b.Buffer.Data[0][b.w-1] = tr
	b.Buffer.Data[b.h-1][0] = bl
	b.Buffer.Data[b.h-1][b.w-1] = br

	textW := b.w - 2
	textH := 1

	for i, let := range b.text {
		if i >= textW {
			break
		}
		b.Buffer.Data[textH][i+1] = let
	}
}
