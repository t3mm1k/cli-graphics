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
		clipH:   min(h, max(0, c.clipH-y)),
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

func (c *Canvas) DrawString(x, y int, s string, fgColor Color) {
	r := []rune(s)

	for i, ch := range r {
		globalX := c.offsetX + x + i
		globalY := c.offsetY + y

		bgColor := ColorDefault()
		if globalX >= 0 && globalX < c.buffer.W && globalY >= 0 && globalY < c.buffer.H {
			bgColor = c.buffer.Data[globalY][globalX].BgColor
		}

		c.SetCell(x+i, y, NewCellColored(ch, fgColor, bgColor))
	}
}

func (c *Canvas) DrawRect(x, y, w, h int, borderStyle BorderStyle, borderColor Color, bgColor Color) {
	if w < 2 || h < 2 {
		return
	}

	// Заливаем фон внутренности (без обводки)
	// Вариант А: фон только внутри, обводка с дефолтным цветом
	// Вариант Б: фон включая обводку
	// Выбран Вариант А — обводка остаётся с borderColor, фон только внутри
	if !bgColor.IsDefault {
		c.Fill(x+1, y+1, w-2, h-2, NewCellColored(' ', ColorDefault(), bgColor))
	}

	var hLine, vLine Cell
	var tl, tr, bl, br Cell

	switch borderStyle {
	case BorderDouble:
		hLine, vLine = NewCellColored('═', borderColor, ColorDefault()), NewCellColored('║', borderColor, ColorDefault())
		tl = NewCellColored('╔', borderColor, ColorDefault())
		tr = NewCellColored('╗', borderColor, ColorDefault())
		bl = NewCellColored('╚', borderColor, ColorDefault())
		br = NewCellColored('╝', borderColor, ColorDefault())
	case BorderRounded:
		hLine, vLine = NewCellColored('─', borderColor, ColorDefault()), NewCellColored('│', borderColor, ColorDefault())
		tl = NewCellColored('╭', borderColor, ColorDefault())
		tr = NewCellColored('╮', borderColor, ColorDefault())
		bl = NewCellColored('╰', borderColor, ColorDefault())
		br = NewCellColored('╯', borderColor, ColorDefault())
	default: // BorderSingle
		hLine, vLine = NewCellColored('─', borderColor, ColorDefault()), NewCellColored('│', borderColor, ColorDefault())
		tl = NewCellColored('┌', borderColor, ColorDefault())
		tr = NewCellColored('┐', borderColor, ColorDefault())
		bl = NewCellColored('└', borderColor, ColorDefault())
		br = NewCellColored('┘', borderColor, ColorDefault())
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

func (c *Canvas) Fill(x, y, w, h int, cell Cell) {
	for row := y; row < y+h; row++ {
		for col := x; col < x+w; col++ {
			c.SetCell(col, row, cell)
		}
	}
}