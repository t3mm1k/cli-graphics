package engine

import (
	"fmt"
	"strings"
)

type Buffer struct {
	Data [][]rune
	W, H int
}

func (b *Buffer) GetWidth() int {
	return b.W
}

func (b *Buffer) GetHeight() int {
	return b.H
}

func (b *Buffer) GetObjects() [][]rune {
	return b.Data
}

func NewBuffer(width, height int) *Buffer {

	var buf = make([][]rune, height)
	for i := range height {
		buf[i] = make([]rune, width)

		for j := range width {
			buf[i][j] = ' '
		}
	}

	return &Buffer{
		H:    height,
		W:    width,
		Data: buf,
	}
}

func (b *Buffer) Clear() {
	for i := range b.Data {
		for j := range b.Data[i] {
			b.Data[i][j] = ' '
		}
	}
}

func (b *Buffer) Blit(child [][]rune, x, y int) {
	for i, bufferLine := range child {
		copy(b.Data[y+i][x:], bufferLine)
	}
}

func (b *Buffer) Flush() {
	var builder strings.Builder

	builder.WriteString("\033[H")

	for _, row := range b.Data {
		builder.WriteString(string(row))
		builder.WriteString("\n")
	}

	fmt.Print(builder.String())
}
