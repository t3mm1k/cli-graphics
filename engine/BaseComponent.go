package engine

type BaseComponent struct {
	x, y   int
	w, h   int
	Buffer *Buffer
	Border bool
}

func NewBaseComponent(x, y, w, h int, border bool) BaseComponent {
	return BaseComponent{
		x: x, y: y, w: w, h: h,
		Buffer: NewBuffer(w, h),
		Border: border,
	}
}

func (b *BaseComponent) RenderBorder(isFocused bool) {
	if !b.Border {
		return
	}

	var hLine, vLine rune
	var tl, tr, bl, br rune

	if isFocused {
		hLine, vLine = '═', '║'
		tl, tr, bl, br = '╔', '╗', '╚', '╝'
	} else {
		hLine, vLine = '─', '│'
		tl, tr, bl, br = '┌', '┐', '└', '┘'
	}

	for x := 1; x < b.w-1; x++ {
		b.Buffer.Data[0][x] = hLine
		b.Buffer.Data[b.h-1][x] = hLine
	}

	for y := 1; y < b.h-1; y++ {
		b.Buffer.Data[y][0] = vLine
		b.Buffer.Data[y][b.w-1] = vLine
	}

	b.Buffer.Data[0][0] = tl
	b.Buffer.Data[0][b.w-1] = tr
	b.Buffer.Data[b.h-1][0] = bl
	b.Buffer.Data[b.h-1][b.w-1] = br
}

func (b *BaseComponent) GetBuffer() [][]rune {
	return b.Buffer.GetObjects()
}

func (b *BaseComponent) GetSize() (w, h int) {
	return b.w, b.h
}

func (b *BaseComponent) SetSize(w, h int) {
	b.w, b.h = w, h
}

func (b *BaseComponent) GetCoords() (x, y int) {
	return b.x, b.y
}
