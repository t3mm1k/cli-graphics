package widgets

import (
	"cli-graphics/engine"
	"fmt"
)

type ListView struct {
	W, H, X, Y int

	Lines []string

	Buffer *engine.Buffer
}

func NewList(w, h, x, y int, lines []string) *ListView {
	buf := engine.NewBuffer(w, h)

	return &ListView{
		W:      w,
		H:      h,
		X:      x,
		Y:      y,
		Lines:  lines,
		Buffer: buf,
	}
}

func (l *ListView) Coords() (x, y int) {
	return l.X, l.Y
}

func (l *ListView) GetBuffer() [][]rune {
	return l.Buffer.GetObjects()
}

func (l *ListView) Render() {
	for x := 1; x < l.W-1; x++ {
		l.Buffer.Data[0][x] = '─'
		l.Buffer.Data[l.H-1][x] = '─'
	}

	for y := 1; y < l.H-1; y++ {
		l.Buffer.Data[y][0] = '│'
		l.Buffer.Data[y][l.W-1] = '│'
	}

	l.Buffer.Data[0][0] = '┌'
	l.Buffer.Data[0][l.W-1] = '┐'
	l.Buffer.Data[l.H-1][0] = '└'
	l.Buffer.Data[l.H-1][l.W-1] = '┘'

	for i, line := range l.Lines {
		if i >= l.H-2 {
			break
		}

		source := []rune("• " + line)
		maxBufferAvailable := l.W - 2

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
