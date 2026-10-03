package widgets

import (
	"cli-graphics/engine"
	"cli-graphics/utils"
	"unicode/utf8"

	"github.com/google/uuid"
)

type Label struct {
	engine.BaseComponent
	text string
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
	}

	defaultClasses := "border-none fg:default text-left"
	if border {
		defaultClasses = "border-single border:dim-gray fg:default text-left"
	}
	lbl.InitStyle(defaultClasses)

	engine.Registry.AddComponent(lbl)

	return lbl
}

func (l *Label) Render(canvas *engine.Canvas) {
	style := l.CurrentStyle()
	w, h := l.GetSize()

	offset := 0
	if style.Border != engine.BorderNone {
		canvas.DrawRect(0, 0, w, h, style.Border, style.BorderFg, style.Bg)
		offset = 1
	}

	textX := offset + style.Padding.Left
	switch style.Align {
	case engine.TextAlignCenter:
		textX = w / 2
	case engine.TextAlignRight:
		textX = w - offset - style.Padding.Right
	}
	textY := offset + style.Padding.Top

	canvas.DrawStringAligned(textX, textY, l.text, style.Fg, style.TextStyle, style.Align)
}

func (l *Label) OnTick() {}
