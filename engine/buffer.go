package engine

import (
	"fmt"
	"strings"
)

type Buffer struct {
	Data [][]Cell
	W, H int
}

func (b *Buffer) GetSize() (int, int) {
	return b.W, b.H
}

func (b *Buffer) SetSize(w, h int) {
	if w <= 0 || h <= 0 {
		panic(&InvalidBufferSizeError{W: w, H: h})
	}
	if b.W == w && b.H == h {
		return
	}
	newData := make([][]Cell, h)
	for i := range h {
		newData[i] = make([]Cell, w)
		for j := range w {
			newData[i][j] = NewCell(' ')
		}
	}
	b.W = w
	b.H = h
	b.Data = newData
}

func (b *Buffer) GetObjects() [][]Cell {
	return b.Data
}

func NewBuffer(width, height int) (*Buffer, error) {
	if (width <= 0) || (height <= 0) {
		return nil, &InvalidBufferSizeError{W: width, H: height}
	}

	var buf = make([][]Cell, height)
	for i := range height {
		buf[i] = make([]Cell, width)

		for j := range width {
			buf[i][j] = NewCell(' ')
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
			b.Data[i][j] = NewCell(' ')
		}
	}
}

func (b *Buffer) ClearRegion(x, y, w, h int) {
	for row := y; row < y+h && row < b.H; row++ {
		for col := x; col < x+w && col < b.W; col++ {
			b.Data[row][col] = NewCell(' ')
		}
	}
}

func (b *Buffer) Flush() string {
    var builder strings.Builder

    lastFg := ColorDefault()
    lastBg := ColorDefault()
    lastStyle := TextStyleDefault

    for y, row := range b.Data {
		for _, cell := range row {
			if cell.BgColor != lastBg {
				if cell.BgColor.IsDefault {
					builder.WriteString("\033[49m")
				} else {
					fmt.Fprintf(&builder, "\033[48;2;%d;%d;%dm", cell.BgColor.R, cell.BgColor.G, cell.BgColor.B)
				}
				lastBg = cell.BgColor
			}

			if cell.FgColor != lastFg {
				if cell.FgColor.IsDefault {
					builder.WriteString("\033[39m")
				} else {
					fmt.Fprintf(&builder, "\033[38;2;%d;%d;%dm", cell.FgColor.R, cell.FgColor.G, cell.FgColor.B)
				}
				lastFg = cell.FgColor
			}

			if cell.Style != lastStyle {
				switch cell.Style {
				case TextStyleDefault:
					builder.WriteString("\033[0m")
				case TextStyleBold:
					builder.WriteString("\033[1m")
				case TextStyleDim:
					builder.WriteString("\033[2m")
				case TextStyleItalic:
					builder.WriteString("\033[3m")
				case TextStyleUnderline:
					builder.WriteString("\033[4m")
				case TextStyleBlink:
					builder.WriteString("\033[5m")
				case TextStyleReverse:
					builder.WriteString("\033[7m")
				case TextStyleHidden:
					builder.WriteString("\033[8m")
				case TextStyleStrikethrough:
					builder.WriteString("\033[9m")
				}

				lastStyle = cell.Style
			}


			builder.WriteRune(cell.R)
		}

		if y < len(b.Data)-1 {
			builder.WriteString("\r\n")
		}
	}
	builder.WriteString("\033[0m")
	return builder.String()
}
