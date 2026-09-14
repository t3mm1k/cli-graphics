package widgets

import (
	"cli-graphics/engine"
	"cli-graphics/utils"
	"unicode/utf8"

	"github.com/google/uuid"
)

type Label struct {
	text string

	engine.BaseComponent
}

func NewLabel(x, y int, border bool, text string, l ...int) *Label {
	var width int

	if len(l) > 0 {
		width = l[0]
	} else {
		width = utf8.RuneCountInString(text)
	}
	width, height := utils.GetActuallySize(border, width, 1)

	id := uuid.New()
	lbl := &Label{
		BaseComponent: engine.NewBaseComponent(id, x, y, width, height, border),
		text:          text,
	}

	return lbl
}

func (l *Label) Render() {

	l.RenderBorder(false)

	w, _ := l.GetSize()

	var textW, textH int = w, 0

	if l.Border {
		textH = 1
		textW -= 2
	}

	text := []rune(l.text)

	for i, let := range text {
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

func (l *Label) OnTick() {}

func (l *Label) GetBuffer() [][]rune {
	return l.Buffer.GetObjects()
}
