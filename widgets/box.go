package widgets

import "cli-graphics/engine"

type Box struct {
	W, H int
	X, Y int

	Buffer *engine.Buffer

	Children []engine.Component
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
