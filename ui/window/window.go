package window

import (
	"cli-graphics/ui/listView"
	"fmt"
)

type Window struct {
	Length int
	Width  int

	Buffer  [][]rune
	Objects []*listView.ListView
}

func (w *Window) clearBuffer() {
	for i := 0; i < w.Length; i++ {
		for j := 0; j < w.Width; j++ {
			w.Buffer[i][j] = ' '
		}
	}
}

func (w *Window) Render() {
	for i := 1; i < w.Width-1; i++ {
		w.Buffer[0][i] = '─'
		w.Buffer[w.Length-1][i] = '─'
	}

	for i := 1; i < w.Length-1; i++ {
		w.Buffer[i][0] = '│'
		w.Buffer[i][w.Width-1] = '│'
	}

	w.Buffer[0][0] = '┌'
	w.Buffer[0][w.Width-1] = '┐'
	w.Buffer[w.Length-1][0] = '└'
	w.Buffer[w.Length-1][w.Width-1] = '┘'

	for _, obj := range w.Objects {
		obj.Render()
		var objectBuffer = obj.Buffer
		var x, y = obj.Coords.X, obj.Coords.Y
		y = y - 1
		for i, bufferLine := range objectBuffer {
			copy(w.Buffer[y+i][x:], bufferLine)
		}
	}
}

func (w *Window) AddObject(object *listView.ListView) {
	w.Objects = append(w.Objects, object)
}

func (w *Window) Print() {
	for i := 0; i < w.Length; i++ {
		fmt.Println(string(w.Buffer[i]))
	}
}

func New(length int, width int) *Window {
	buf := make([][]rune, length)

	for i := 0; i < length; i++ {
		buf[i] = make([]rune, width)
		for j := 0; j < width; j++ {
			buf[i][j] = ' '
		}
	}

	return &Window{
		Width:  width,
		Length: length,
		Buffer: buf,
	}
}
