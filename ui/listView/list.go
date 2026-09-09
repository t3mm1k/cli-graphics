package listView

import (
	"cli-graphics/ui"
	"fmt"
)

type ListView struct {
	MaxLength int
	MaxWidth  int

	Coords ui.Point // координата левого верхнего угла рендер начнется со следующей клетки по горизонтали и с той же по вертикали
	Lines  []string

	Buffer [][]rune
}

func New(MaxLength int, MaxWidth int, Coords ui.Point, Lines []string) *ListView {
	buf := make([][]rune, MaxLength)

	for i := 0; i < MaxLength; i++ {
		buf[i] = make([]rune, MaxWidth)
		for j := 0; j < MaxWidth; j++ {
			buf[i][j] = ' '
		}
	}

	return &ListView{
		MaxLength: MaxLength,
		MaxWidth:  MaxWidth,
		Coords:    Coords,
		Lines:     Lines,
		Buffer:    buf,
	}
}

func (l *ListView) Render() {
	for i := 1; i < l.MaxWidth-1; i++ {
		l.Buffer[0][i] = '─'
		l.Buffer[l.MaxLength-1][i] = '─'
	}

	for i := 1; i < l.MaxLength-1; i++ {
		l.Buffer[i][0] = '│'
		l.Buffer[i][l.MaxWidth-1] = '│'
	}

	l.Buffer[0][0] = '┌'
	l.Buffer[0][l.MaxWidth-1] = '┐'
	l.Buffer[l.MaxLength-1][0] = '└'
	l.Buffer[l.MaxLength-1][l.MaxWidth-1] = '┘'

	for i, line := range l.Lines {
		if i >= l.MaxLength {
			break
		}

		source := []rune("• " + line)
		maxBufferAvailable := l.MaxWidth - 2

		if len(source) > maxBufferAvailable {
			source = source[:maxBufferAvailable]
		}

		copy(l.Buffer[i+1][1:], source)
	}
}

func (l *ListView) Print() {
	for _, row := range l.Buffer {
		fmt.Println(string(row))
	}
}
