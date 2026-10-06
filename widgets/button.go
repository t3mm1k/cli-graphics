package widgets

import (
	"cli-graphics/engine"
	"cli-graphics/utils"
	"unicode/utf8"

	"github.com/google/uuid"
)

const DefaultButtonClasses = "border-single border:dim-gray fg:default text-center focus:border-double focus:border:default disabled:dim active:reverse hover:border:white"

type Button struct {
	engine.BaseComponent

	text    string
	OnClick func()
}

func NewButton(x, y int, text string, onClick func(), classes ...string) *Button {
	id := uuid.New()
	width := utf8.RuneCountInString(text)
	width, height := utils.GetActuallySize(true, width, 1)
	btn := &Button{
		BaseComponent: engine.NewBaseComponent(id, x, y, width, height),
		text:          text,
		OnClick:       onClick,
	}

	btn.InitStyle(DefaultButtonClasses, classes...)

	engine.Registry.AddComponent(btn)
	engine.FocusManagerInstance.Register(id)

	return btn
}

func (b *Button) SetText(newText string) {
	b.text = newText
	width := utf8.RuneCountInString(b.text)
	width, height := utils.GetActuallySize(true, width, 1)
	b.SetSize(width, height)
}

func (b *Button) SetFocus(focused bool) {
	// FocusManager сам отслеживает фокус
}

func (b *Button) IsFocused() bool {
	return engine.FocusManagerInstance.GetFocused() == b.GetId()
}

func (b *Button) OnTick() {}

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
	style := b.CurrentStyle()
	w, h := b.GetSize()

	canvas.DrawRect(0, 0, w, h, style.Border, style.BorderFg, style.Bg)

	offset := 1
	if style.Border == engine.BorderNone {
		offset = 0
	}
	
	textX := offset + style.Padding.Left
	switch style.Align {
	case engine.TextAlignCenter:
		textX = offset + style.Padding.Left + (w-offset*2-style.Padding.Left-style.Padding.Right)/2
	case engine.TextAlignRight:
		textX = w - offset - style.Padding.Right
	}
	textY := offset + style.Padding.Top

	canvas.DrawStringAligned(textX, textY, b.text, style.Fg, style.TextStyle, style.Align)
}


func (b *Button) SetColors(normal, focus engine.Color) {
	b.Style.BorderFg = normal
	b.FocusedStyle.BorderFg = focus
}

func (b *Button) HandleMouse(e *engine.MouseEvent) bool {
	if e.MouseButton == utils.MouseBtnLeft {
		if b.OnClick != nil {
			b.OnClick()
		}
		return true
	}
	return false
}
