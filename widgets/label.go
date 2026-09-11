package widgets

import (
	"cli-graphics/engine"
	"unicode/utf8"
)

type Label struct {
	x, y int

	text string

	w, h int

	Border bool

	Buffer *engine.Buffer
}

func GetActuallySize(hasBorder bool, w, h int) (W, H int) {
	if hasBorder {
		return w + 2, h + 2
	}
	return w, h
}

func NewLabel(x, y int, border bool, text string, l ...int) *Label {
	var width int

	if len(l) > 0 {
		width = l[0]
	} else {
		width = utf8.RuneCountInString(text)
	}

	width, height := GetActuallySize(border, width, 1)

	buf := engine.NewBuffer(width, height)

	return &Label{
		text:   text,
		x:      x,
		y:      y,
		w:      width,
		h:      height,
		Border: border,
		Buffer: buf,
	}
}

func (l *Label) GetCoords() (x, y int) {
	return l.x, l.y
}

func (l *Label) GetSize() (w, h int) {
	return l.w, l.h
}

func (l *Label) Render() {

	if l.Border {
		for x := 1; x < l.w-1; x++ {
			l.Buffer.Data[0][x] = '─'
			l.Buffer.Data[l.h-1][x] = '─'
		}

		for y := 1; y < l.h-1; y++ {
			l.Buffer.Data[y][0] = '│'
			l.Buffer.Data[y][l.w-1] = '│'
		}

		l.Buffer.Data[0][0] = '┌'
		l.Buffer.Data[0][l.w-1] = '┐'
		l.Buffer.Data[l.h-1][0] = '└'
		l.Buffer.Data[l.h-1][l.w-1] = '┘'
	}

	var textW, textH int = l.w, 0

	if l.Border {
		textH = 1
		textW -= 2
	}

	for i, let := range l.text {
		pos := i
		if l.Border {
			pos += 1
		}
		if i >= textW {
			break
		}
		l.Buffer.Data[textH][pos] = let
	}
}

func (l *Label) GetBuffer() [][]rune {
	return l.Buffer.GetObjects()
}
