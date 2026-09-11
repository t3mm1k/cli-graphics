package widgets

//
//import "cli-graphics/engine"
//
//type Input struct {
//	x, y int
//
//	w int
//	h int
//
//	value string
//
//	Buffer *engine.Buffer
//
//	isFocused bool
//
//	OnInput func()
//}
//
//func (i *Input) SetFocus(focused bool) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (i *Input) IsFocused() bool {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (i *Input) HandleKey(key string) bool {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (i *Input) Render() {
//	var hLine, vLine rune
//	var tl, tr, bl, br rune
//
//	if i.isFocused {
//		hLine, vLine = '═', '║'
//		tl, tr, bl, br = '╔', '╗', '╚', '╝'
//	} else {
//		hLine, vLine = '─', '│'
//		tl, tr, bl, br = '┌', '┐', '└', '┘'
//	}
//
//	for x := 1; x < b.W-1; x++ {
//		b.Buffer.Data[0][x] = hLine
//		b.Buffer.Data[b.H-1][x] = hLine
//	}
//
//	for y := 1; y < b.H-1; y++ {
//		b.Buffer.Data[y][0] = vLine
//		b.Buffer.Data[y][b.W-1] = vLine
//	}
//
//	b.Buffer.Data[0][0] = tl
//	b.Buffer.Data[0][b.W-1] = tr
//	b.Buffer.Data[b.H-1][0] = bl
//	b.Buffer.Data[b.H-1][b.W-1] = br
//}
//
//func (i *Input) GetBuffer() [][]rune {
//	return i.Buffer.GetObjects()
//}
//
//func (i Input) GetCoords() (x, y int) {
//	return i.x, i.y
//}
//
//func NewInput(x, y, l int) *Input {
//	buf := engine.NewBuffer(l+2, 3)
//
//	return &Input{
//		x:      x,
//		y:      y,
//		Buffer: buf,
//	}
//}
