package widgets

import (
	"cli-graphics/engine"
	"unicode/utf8"
)

type Label struct {
	X, Y int

	text string

	L int

	Buffer *engine.Buffer
}

func NewLabel(x, y int, text string, l ...int) *Label {
	var length int

	if len(l) > 0 {
		length = l[0]
	} else {
		length = utf8.RuneCountInString(text)
	}

	buf := engine.NewBuffer(length, 1)

	return &Label{
		text:   text,
		X:      x,
		Y:      y,
		L:      length,
		Buffer: buf,
	}
}

func (l *Label) Coords() (x, y int) {
	return l.X, l.Y
}

func (l *Label) Render() {
	for i, let := range l.text {
		if i >= l.L {
			break
		}
		l.Buffer.Data[0][i] = let
	}
}

func (l *Label) GetBuffer() [][]rune {
	return l.Buffer.GetObjects()
}
