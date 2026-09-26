package engine

import "github.com/google/uuid"

type Component interface {
	Render(canvas *Canvas)

	GetCoords() (x, y int)

	GetSize() (w, h int)

	GetId() uuid.UUID

	OnTick()
}
