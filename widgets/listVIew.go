package widgets

import (
	"cli-graphics/engine"
	"fmt"

	"github.com/google/uuid"
)

type ListView struct {
	engine.BaseComponent

	Lines []string
}

func NewList(w, h, x, y int, lines []string) *ListView {
	id := uuid.New()
	list := &ListView{
		BaseComponent: engine.NewBaseComponent(id, x, y, w, h, true),
		Lines:         lines,
	}

	engine.Registry.AddComponent(list)

	return list
}

func (l *ListView) GetBuffer() [][]rune {
	return l.Buffer.GetObjects()
}

func (l *ListView) Render() {
	l.Buffer.Clear()
	l.RenderBorder(false)

	w, h := l.GetSize()

	for i, line := range l.Lines {
		if i >= h-2 {
			break
		}

		source := []rune("• " + line)
		maxBufferAvailable := w - 2

		if len(source) > maxBufferAvailable {
			source = source[:maxBufferAvailable]
		}

		copy(l.Buffer.Data[i+1][1:], source)
		//l.Print()
	}
}

func (l *ListView) OnTick() {}

func (l *ListView) Print() {
	buf := l.Buffer.GetObjects()
	for _, row := range buf {
		fmt.Println(string(row))
	}
}
