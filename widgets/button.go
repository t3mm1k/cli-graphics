package widgets

import (
	"cli-graphics/engine"
	"cli-graphics/utils"
	"unicode/utf8"

	"github.com/google/uuid"
)

type Button struct {
	engine.BaseComponent

	text string

	isFocused bool
	OnClick   func()
}

func NewButton(x, y int, text string, onClick func()) *Button {
	id := uuid.New()
	width := utf8.RuneCountInString(text)
	width, height := utils.GetActuallySize(true, width, 1)
	btn := &Button{
		BaseComponent: engine.NewBaseComponent(id, x, y, width, height),
		text:          text,
		isFocused:     false,
		OnClick:       onClick,
	}

	engine.Registry.AddComponent(btn)

	return btn
}

func (b *Button) SetText(newText string) {
	b.text = newText
	width := utf8.RuneCountInString(b.text)
	width, height := utils.GetActuallySize(true, width, 1)
	b.SetSize(width, height)
}

func (b *Button) SetFocus(focused bool) {
	b.isFocused = focused
}

func (b *Button) OnTick() {}

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

func (b *Button) Render(canvas *engine.Canvas) {
	w, h := b.GetSize()

	borderStyle := engine.BorderSingle
	if b.isFocused {
		borderStyle = engine.BorderDouble
	}

	canvas.DrawRect(0, 0, w, h, borderStyle)

	canvas.DrawString(1, 1, b.text)
}
