package engine

type BorderStyle string

const (
	BorderSingle  BorderStyle = "single"
	BorderDouble  BorderStyle = "double"
	BorderRounded BorderStyle = "rounded"
)

type Canvas struct {
	buffer  *Buffer
	offsetX int
	offsetY int
	clipW   int
	clipH   int
}

func NewCanvas(buffer *Buffer) *Canvas {
	return &Canvas{
		buffer: buffer,
		clipH:  buffer.H,
		clipW:  buffer.W,
	}
}

func (c *Canvas) SubCanvas(x, y, w, h int) *Canvas {
	return &Canvas{
		buffer:  c.buffer,
		offsetX: c.offsetX + x,
		offsetY: c.offsetY + y,
		clipW:   min(w, max(0, c.clipW-x)),
		clipH:   min(h, c.clipH-y),
	}
}

func (c *Canvas) SetCell(x, y int, cell Cell) {
	if (y < 0) || (x < 0) || (x >= c.clipW) || (y >= c.clipH) {
		return
	}

	globalX := c.offsetX + x
	globalY := c.offsetY + y

	if globalX < 0 || globalX >= c.buffer.W || globalY < 0 || globalY >= c.buffer.H {
		return
	}

	c.buffer.Data[globalY][globalX] = cell
}

func (c *Canvas) SetRune(x, y int, rune rune) {
	c.SetCell(x, y, NewCell(rune))
}

func (c *Canvas) DrawString(x, y int, s string) {
	r := []rune(s)

	for i, cell := range r {
		c.SetCell(x+i, y, NewCellColored(cell, ColorNeonCyan, ColorDefault())) //TODO Сделать СТИЛИ ДЛЯ СТРОК!!!
	}
}

func (c *Canvas) DrawRect(x, y, w, h int, borderStyle ...BorderStyle) { //TODO добавить стили так же
	if w < 2 || h < 2 {
		return
	}

	style := BorderSingle
	if len(borderStyle) > 0 {
		style = borderStyle[0]
	}
	var hLine, vLine Cell
	var tl, tr, bl, br Cell

	switch style {
	case BorderDouble:
		hLine, vLine = NewCell('═'), NewCell('║')
		tl, tr, bl, br = NewCell('╔'), NewCell('╗'), NewCell('╚'), NewCell('╝')
	case BorderRounded:
		hLine, vLine = NewCell('─'), NewCell('│')
		tl, tr, bl, br = NewCell('╭'), NewCell('╮'), NewCell('╰'), NewCell('╯')
	default:
		hLine, vLine = NewCell('─'), NewCell('│')
		tl, tr, bl, br = NewCell('┌'), NewCell('┐'), NewCell('└'), NewCell('┘')
	}

	for col := x + 1; col < x+w-1; col++ {
		c.SetCell(col, y, hLine)
		c.SetCell(col, y+h-1, hLine)
	}

	for row := y + 1; row < y+h-1; row++ {
		c.SetCell(x, row, vLine)
		c.SetCell(x+w-1, row, vLine)
	}

	c.SetCell(x, y, tl)
	c.SetCell(x+w-1, y, tr)
	c.SetCell(x, y+h-1, bl)
	c.SetCell(x+w-1, y+h-1, br)
}
