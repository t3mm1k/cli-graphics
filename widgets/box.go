package widgets

import (
	"cli-graphics/engine"

	"github.com/google/uuid"
)

type Box struct {
	engine.BaseComponent
	children []engine.Component
	focused  bool
}

func (b *Box) IsFocused() bool {
	return b.focused
}

func NewBox(x, y, w, h int) *Box {
	id := uuid.New()

	box := &Box{
		BaseComponent: engine.NewBaseComponent(id, x, y, w, h),
	}

	engine.Registry.AddComponent(box)
	return box
}

func (b *Box) Render(canvas *engine.Canvas) {
	w, h := b.GetSize()
	canvas.DrawRect(0, 0, w, h, engine.BorderSingle)

	for _, child := range b.children {
		x, y := child.GetCoords()
		w, h := child.GetSize()
		childCanvas := canvas.SubCanvas(x, y, w, h)
		child.Render(childCanvas)
	}
}

func (b *Box) OnTick() {
	for _, child := range b.children {
		child.OnTick()
	}
}

func (b *Box) AddChild(child engine.Component) {
	if child == nil {
		engine.Log.Warn("AddChild: attempt to add nil child, ignored", "err", engine.NilChildError)
		return
	}
	if child.GetId() == b.GetId() {
		panic(engine.CyclicDependencyError)
	}

	b.children = append(b.children, child)
	engine.Registry.SetParent(child.GetId(), b.GetId())
}

func (b *Box) SetFocus(focused bool) {
	b.focused = focused
}

func (b *Box) HandleKey(key string) bool {
	if key == "Tab" {
		engine.FocusManagerInstance.FocusNext()
		return true
	}
	if key == "Shift+Tab" {
		engine.FocusManagerInstance.FocusPrev()
		return true
	}

	focusedId := engine.FocusManagerInstance.GetFocused()
	if focusedId == uuid.Nil {
		return false
	}

	comp, exists := engine.Registry.GetComponent(focusedId)
	if !exists {
		return false
	}

	if f, ok := comp.(engine.Focusable); ok {
		return f.HandleKey(key)
	}

	return false
}