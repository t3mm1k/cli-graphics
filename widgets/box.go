package widgets

import "cli-graphics/engine"

type Box struct {
	W, H int
	X, Y int

	Buffer *engine.Buffer

	Children []engine.Component

	focusedIndex int
	focused      bool
}

func (b *Box) IsFocused() bool {
	return b.focused
}

func NewBox(x, y, w, h int) *Box {
	return &Box{
		W:      w,
		H:      h,
		X:      x,
		Y:      y,
		Buffer: engine.NewBuffer(w, h),
	}
}

func (b *Box) GetBuffer() [][]rune {
	return b.Buffer.GetObjects()
}

func (b *Box) Coords() (x, y int) {
	return b.X, b.Y
}

func (b *Box) Render() {
	b.Buffer.Clear()

	for x := 1; x < b.W-1; x++ {
		b.Buffer.Data[0][x] = '─'
		b.Buffer.Data[b.H-1][x] = '─'
	}

	for y := 1; y < b.H-1; y++ {
		b.Buffer.Data[y][0] = '│'
		b.Buffer.Data[y][b.W-1] = '│'
	}

	b.Buffer.Data[0][0] = '┌'
	b.Buffer.Data[0][b.W-1] = '┐'
	b.Buffer.Data[b.H-1][0] = '└'
	b.Buffer.Data[b.H-1][b.W-1] = '┘'

	for _, child := range b.Children {
		child.Render()
		x, y := child.Coords()
		b.Buffer.Blit(child.GetBuffer(), x, y)
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
