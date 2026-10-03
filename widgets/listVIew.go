package widgets

import (
	"cli-graphics/engine"

	"github.com/google/uuid"
)

const DefaultListViewClasses = "border-single border:dim-gray fg:default text-left focus:border:neon-pink selected:fg:yellow selected:bold"

type ListView struct {
	engine.BaseComponent

	lines         []string
	selectedIndex int
	focused       bool
}

func NewList(w, h, x, y int, lines []string, classes ...string) *ListView {
	id := uuid.New()
	list := &ListView{
		BaseComponent: engine.NewBaseComponent(id, x, y, w, h),
		lines:         lines,
		selectedIndex: 0,
		focused:       false,
	}

	list.InitStyle(DefaultListViewClasses, classes...)

	engine.Registry.AddComponent(list)
	engine.FocusManagerInstance.Register(id)

	return list
}

func (l *ListView) SetFocus(focused bool) {
	l.focused = focused
}

func (l *ListView) IsFocused() bool {
	return engine.FocusManagerInstance.GetFocused() == l.GetId()
}

func (l *ListView) HandleKey(key string) bool {
	if len(l.lines) == 0 {
		return false
	}

	switch key {
	case "Down", "ArrowDown":
		if l.selectedIndex < len(l.lines)-1 {
			l.selectedIndex++
		} else {
			l.selectedIndex = 0 // зацикливание вниз
		}
		return true

	case "Up", "ArrowUp":
		if l.selectedIndex > 0 {
			l.selectedIndex--
		} else {
			l.selectedIndex = len(l.lines) - 1 // зацикливание вверх
		}
		return true
	}

	return false
}

func (l *ListView) Render(canvas *engine.Canvas) {
	style := l.CurrentStyle()
	w, h := l.GetSize()

	offset := 0
	if style.Border != engine.BorderNone {
		canvas.DrawRect(0, 0, w, h, style.Border, style.BorderFg, style.Bg)
		offset = 1
	}

	for i, line := range l.lines {
		if offset > 0 && i >= h-2 {
			break
		} else if offset == 0 && i >= h {
			break
		}

		source := "• " + line
		textX := offset + style.Padding.Left
		textY := offset + style.Padding.Top + i

		canvas.DrawStringAligned(textX, textY, source, style.Fg, style.TextStyle, style.Align)
	}
}

func (l *ListView) OnTick() {}
