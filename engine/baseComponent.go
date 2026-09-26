package engine

import "github.com/google/uuid"

type BaseComponent struct {
	id   uuid.UUID
	x, y int
	w, h int
}

func NewBaseComponent(id uuid.UUID, x, y, w, h int) BaseComponent {
	return BaseComponent{
		id: id,
		x:  x, y: y, w: w, h: h,
	}
}

func (b *BaseComponent) GetSize() (w, h int) {
	return b.w, b.h
}

func (b *BaseComponent) SetSize(w, h int) {
	b.w, b.h = w, h
}

func (b *BaseComponent) GetId() uuid.UUID {
	return b.id
}

func (b *BaseComponent) GetCoords() (x, y int) {
	return b.x, b.y
}
