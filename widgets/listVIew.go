package widgets

import (
	"cli-graphics/engine"

	"github.com/google/uuid"
)

type ListView struct {
	engine.BaseComponent

	lines []string
}

func NewList(w, h, x, y int, lines []string) *ListView {
	id := uuid.New()
	list := &ListView{
		BaseComponent: engine.NewBaseComponent(id, x, y, w, h),
		lines:         lines,
	}

	engine.Registry.AddComponent(list)

	return list
}

func (l *ListView) Render(canvas *engine.Canvas) {

	_, h := l.GetSize()

	for i, line := range l.lines {
		if i >= h-2 {
			break
		}

		source := "• " + line

		canvas.DrawString(1, i+1, source, engine.ColorDefault())
		//l.Print()
	}
}

func (l *ListView) OnTick() {}
