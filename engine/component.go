package engine

type Component interface {
	Render()

	GetBuffer() [][]rune

	Coords() (x, y int)
}
