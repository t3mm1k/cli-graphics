package widgets

import (
	"cli-graphics/engine"
	"unicode/utf8"
)

type Label struct {
	X, Y int

	text string

	W int

	Border bool

	Buffer *engine.Buffer
}

// TODO родить функцию наверное не привязанную к классу
func (l *Label) GetActuallySize() (w, h int) {
	if l.Border {
		return l.W + 2, 3
	}
	return l.W, 1
}

// TODO создавать буффер с учетом бордера
func NewLabel(x, y int, text string, l ...int) *Label {
	var width int

	if len(l) > 0 {
		width = l[0]
	} else {
		width = utf8.RuneCountInString(text)
	}

	buf := engine.NewBuffer(width, 1)

	return &Label{
		text:   text,
		X:      x,
		Y:      y,
		W:      width,
		Buffer: buf,
	}
}

func (l *Label) Coords() (x, y int) {
	return l.X, l.Y
}

// TODO рендерить бордер при необходимости
func (l *Label) Render() {
	for i, let := range l.text {
		if i >= l.W {
			break
		}
		l.Buffer.Data[0][i] = let
	}
}

func (l *Label) GetBuffer() [][]rune {
	return l.Buffer.GetObjects()
}
