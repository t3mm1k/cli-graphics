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

func NewLabel(x, y int, border bool, text string, classes ...string) *Label {
	width := utf8.RuneCountInString(text)
	width, height := utils.GetActuallySize(border, width, 1)

	id := uuid.New()
	lbl := &Label{
		BaseComponent: engine.NewBaseComponent(id, x, y, width, height),
		text:          text,
	}

	defaultClasses := "border-none fg:white text-left"
	if border {
		defaultClasses = "border-single border:dim-gray fg:white text-left"
	}
	lbl.InitStyle(defaultClasses, classes...)

	engine.Registry.AddComponent(lbl)

	return lbl
}

func (l *Label) SetText(text string) {
	l.text = text
	width := utf8.RuneCountInString(text)
	border := l.Style.Border != engine.BorderNone
	width, height := utils.GetActuallySize(border, width, 1)
	l.SetSize(width, height)
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
		textX = offset + style.Padding.Left + (w-offset*2-style.Padding.Left-style.Padding.Right)/2
	case engine.TextAlignRight:
		textX = w - offset - style.Padding.Right
	}
	textY := offset + style.Padding.Top

	canvas.DrawStringAligned(textX, textY, l.text, style.Fg, style.TextStyle, style.Align)
}

func (l *Label) OnTick() {}
