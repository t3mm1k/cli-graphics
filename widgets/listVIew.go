package widgets

import (
	"cli-graphics/engine"
	"fmt"
)

type ListView struct {
	w, h, x, y int

	Lines []string

	Buffer *engine.Buffer
}

func NewList(w, h, x, y int, lines []string) *ListView {
	buf := engine.NewBuffer(w, h)

	return &ListView{
		w:      w,
		h:      h,
		x:      x,
		y:      y,
		Lines:  lines,
		Buffer: buf,
	}
}

func (l *ListView) GetCoords() (x, y int) {
	return l.x, l.y
}

func (l *ListView) GetSize() (w, h int) {
	return l.w, l.h
}

func (l *ListView) GetBuffer() [][]rune {
	return l.Buffer.GetObjects()
}

func (l *ListView) Render() {
	for x := 1; x < l.w-1; x++ {
		l.Buffer.Data[0][x] = '─'
		l.Buffer.Data[l.h-1][x] = '─'
	}

	for y := 1; y < l.h-1; y++ {
		l.Buffer.Data[y][0] = '│'
		l.Buffer.Data[y][l.w-1] = '│'
	}

	l.Buffer.Data[0][0] = '┌'
	l.Buffer.Data[0][l.w-1] = '┐'
	l.Buffer.Data[l.w-1][0] = '└'
	l.Buffer.Data[l.h-1][l.w-1] = '┘'

	for i, line := range l.Lines {
		if i >= l.h-2 {
			break
		}

		source := []rune("• " + line)
		maxBufferAvailable := l.w - 2

		if len(source) > maxBufferAvailable {
			source = source[:maxBufferAvailable]
		}

		copy(l.Buffer.Data[i+1][1:], source)
		//l.Print()
	}
}

func (l *ListView) Print() {
	buf := l.Buffer.GetObjects()
	for _, row := range buf {
		fmt.Println(string(row))
	}
}
