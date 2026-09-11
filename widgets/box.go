package widgets

import "cli-graphics/engine"

type Box struct {
	engine.BaseComponent
	Children []engine.Component

	focusedIndex int
	focused      bool
}

func (b *Box) IsFocused() bool {
	return b.focused
}

func NewBox(x, y, w, h int) *Box {
	return &Box{
		BaseComponent: engine.NewBaseComponent(x, y, w, h, true),
	}
}

func (b *Box) GetBuffer() [][]rune {
	return b.Buffer.GetObjects()
}

func (b *Box) Render() {
	b.Buffer.Clear()

	//w, h := b.GetSize()

	b.RenderBorder(false)

	for _, child := range b.Children {
		child.Render()
		x, y := child.GetCoords()
		b.Buffer.Blit(child.GetBuffer(), x, y)
	}
}

func (b *Box) OnTick() {
	for _, child := range b.Children {
		child.OnTick()
	}
}

func (b *Box) AddChild(child engine.Component) {
	b.Children = append(b.Children, child)
}

func (b *Box) FindNextFocusableChild(st int) int {
	for i := st; i < len(b.Children); i++ {
		if _, ok := b.Children[i].(engine.Focusable); ok {
			return i
		}
	}
	return -1
}

func (b *Box) SetFocus(focused bool) {
	b.focused = focused

	if focused {
		b.focusedIndex = b.FindNextFocusableChild(0)

		if b.focusedIndex != -1 {
			if focusable, ok := b.Children[b.focusedIndex].(engine.Focusable); ok {
				focusable.SetFocus(true)
			}
		}
	} else {
		if b.focusedIndex != -1 {
			if focusable, ok := b.Children[b.focusedIndex].(engine.Focusable); ok {
				focusable.SetFocus(false)
			}
		}

		b.focusedIndex = 0
	}
}

func (b *Box) HandleKey(key string) bool {

	if len(b.Children) == 0 || b.focusedIndex == -1 {
		return false
	}

	child := b.Children[b.focusedIndex]
	activeChild, isFocusable := child.(engine.Focusable)

	if isFocusable && activeChild.HandleKey(key) {
		return true
	}

	if key == "Tab" {
		nextIndex := b.FindNextFocusableChild(b.focusedIndex + 1)
		if nextIndex != -1 {
			if isFocusable {
				activeChild.SetFocus(false)
			}
			b.focusedIndex = nextIndex
			if focusable, ok := b.Children[nextIndex].(engine.Focusable); ok {
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
