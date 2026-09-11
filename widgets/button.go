package widgets

import (
	"cli-graphics/engine"
	"cli-graphics/utils"
	"unicode/utf8"
)

type Button struct {
	engine.BaseComponent

	text string

	isFocused bool
	OnClick   func()
}

func NewButton(x, y int, text string, onClick func()) *Button {
	width := utf8.RuneCountInString(text)
	width, height := utils.GetActuallySize(true, width, 1)

	return &Button{
		BaseComponent: engine.NewBaseComponent(x, y, width, height, true),
		text:          text,
		isFocused:     false,
		OnClick:       onClick,
	}
}

func (b *Button) SetText(newText string) {
	b.text = newText

	width := utf8.RuneCountInString(newText)

	b.SetSize(utils.GetActuallySize(true, width, 1))
	b.Buffer = engine.NewBuffer(b.GetSize())
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
	b.RenderBorder(b.isFocused)

	w, _ := b.GetSize()

	textW := w - 2
	textH := 1

	for i, let := range b.text {
		if i >= textW {
			break
		}
		b.Buffer.Data[textH][i+1] = let
	}
}
