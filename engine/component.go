package engine

import "github.com/google/uuid"

type Component interface {
	Render()

	GetBuffer() [][]rune

	GetCoords() (x, y int)

	GetSize() (w, h int)

	GetId() uuid.UUID

	OnTick()

	Rerender()
}
