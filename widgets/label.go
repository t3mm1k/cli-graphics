package widgets

import (
	"cli-graphics/engine"
	"cli-graphics/utils"
	"unicode/utf8"

	"github.com/google/uuid"
)

type Label struct {
	text   string
	border bool
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
		BaseComponent: engine.NewBaseComponent(id, x, y, width, height),
		text:          text,
		border:        border,
	}

	engine.Registry.AddComponent(lbl)

	return lbl
}

func (l *Label) Render(canvas *engine.Canvas) {
	if l.border {
		w, h := l.GetSize()
		canvas.DrawRect(0, 0, w, h, engine.BorderSingle, engine.ColorDefault(), engine.ColorDefault())
		canvas.DrawString(1, 1, l.text, engine.ColorDefault())
	} else {
		canvas.DrawString(0, 0, l.text, engine.ColorDefault())
	}
}

func (l *Label) OnTick() {}
