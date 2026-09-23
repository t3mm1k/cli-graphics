package widgets

import (
	"cli-graphics/engine"

	"github.com/google/uuid"
)

type Box struct {
	engine.BaseComponent
	children []engine.Component

	focusedIndex int
	focused      bool
}

func (b *Box) Flush() {
	if b.Buffer != nil {
		b.Buffer.Flush()
	}
}

func (b *Box) IsFocused() bool {
	return b.focused
}

func NewBox(x, y, w, h int) *Box {
	id := uuid.New()

	box := &Box{
		BaseComponent: engine.NewBaseComponent(id, x, y, w, h, true),
	}

	engine.Registry.AddComponent(box)
	return box
}

func (b *Box) GetBuffer() [][]rune {
	return b.Buffer.GetObjects()
}

func (b *Box) Render() {
	b.Buffer.Clear()

	b.RenderBorder(b.focused)

	for _, child := range b.children {
		child.Render()
		x, y := child.GetCoords()
		b.Buffer.Blit(child.GetBuffer(), x, y)
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

func (b *Box) findNextFocusableChild(st int) int {
	for i := st; i < len(b.children); i++ {
		if _, ok := b.children[i].(engine.Focusable); ok {
			return i
		}
	}
	return -1
}

func (b *Box) SetFocus(focused bool) {
	b.focused = focused

	if focused {
		b.focusedIndex = b.findNextFocusableChild(0)

		if b.focusedIndex != -1 {
			if focusable, ok := b.children[b.focusedIndex].(engine.Focusable); ok {
				focusable.SetFocus(true)
			}
		}
	} else {
		if b.focusedIndex != -1 {
			if focusable, ok := b.children[b.focusedIndex].(engine.Focusable); ok {
				focusable.SetFocus(false)
			}
		}

		b.focusedIndex = 0
	}
}

func (b *Box) HandleKey(key string) bool {

	if len(b.children) == 0 || b.focusedIndex == -1 {
		return false
	}

	child := b.children[b.focusedIndex]
	activeChild, isFocusable := child.(engine.Focusable)

	if isFocusable && activeChild.HandleKey(key) {
		return true
	}

	if key == "Tab" {
		nextIndex := b.findNextFocusableChild(b.focusedIndex + 1)
		if nextIndex != -1 {
			if isFocusable {
				activeChild.SetFocus(false)
			}
			b.focusedIndex = nextIndex
			if focusable, ok := b.children[nextIndex].(engine.Focusable); ok {
				focusable.SetFocus(true)
			}
			return true
		} else {
			if isFocusable {
				activeChild.SetFocus(false)
			}
			b.focusedIndex = -1
			return false
		}
	}
	return false
}
