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

func NewBuffer(width, height int) (*Buffer, error) {
	if (width <= 0) || (height <= 0) {
		return nil, &InvalidBufferSizeError{W: width, H: height}
	}

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
	}, nil
}

func (b *Buffer) Clear() {
	for i := range b.Data {
		for j := range b.Data[i] {
			b.Data[i][j] = ' '
		}
	}
}

func (b *Buffer) ClearRegion(x, y, w, h int) {
	for row := y; row < y+h && row < b.H; row++ {
		for col := x; col < x+w && col < b.W; col++ {
			b.Data[row][col] = ' '
		}
	}
}

func (b *Buffer) Blit(child [][]rune, x, y int) error {
	if b == nil || b.Data == nil || child == nil {
		Log.Error("Blit: target buffer or child data is nil", "err", NilBufferError)
		return NilBufferError
	}
	childH := len(child)
	if childH == 0 {
		return nil
	}
	childW := len(child[0])
	isOutOfBounds := x < 0 || y < 0 || x+childW > b.W || y+childH > b.H
	if isOutOfBounds {
		err := &OutOfBoundsError{
			X: x, Y: y, W: childW, H: childH,
			BufferW: b.W, BufferH: b.H,
		}
		Log.Warn("Blit: child region exceeds bounds, clipping applied", "error", err)
		if x >= b.W || y >= b.H || x+childW <= 0 || y+childH <= 0 {
			return err
		}
	}
	for i, bufferLine := range child {
		destY := y + i
		if destY < 0 || destY >= b.H {
			continue
		}
		lineLen := len(bufferLine)
		destX := x
		srcX := 0
		if destX < 0 {
			srcX = -destX
			destX = 0
		}
		if srcX >= lineLen || destX >= b.W {
			continue
		}
		availableWidth := b.W - destX
		toCopy := lineLen - srcX
		if toCopy > availableWidth {
			toCopy = availableWidth
		}
		copy(b.Data[destY][destX:destX+toCopy], bufferLine[srcX:srcX+toCopy])
	}
	return nil
}

func (b *Buffer) Flush() {
	var builder strings.Builder
	builder.WriteString("\u001B[H\u001B[2J\u001B[3J")

	for _, row := range b.Data {
		builder.WriteString(string(row))
		builder.WriteString("\n")
	}

	fmt.Print(builder.String())
}
