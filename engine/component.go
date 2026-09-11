package engine

type Component interface {
	Render()

	GetBuffer() [][]rune

	GetCoords() (x, y int)

	GetSize() (w, h int)

	OnTick()
}
