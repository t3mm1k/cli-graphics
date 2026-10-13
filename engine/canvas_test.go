package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCanvas_SetCell(t *testing.T) {
	buf, err := NewBuffer(10, 10)
	require.NoError(t, err)
	canvas := NewCanvas(buf)
	ColorRed, _ := BaseColors.Get("red")
	ColorGreen, _ := BaseColors.Get("green")
	testCell := NewCellColored('Z', ColorRed, ColorGreen)

	tests := []struct {
		name        string
		x, y        int
		shouldWrite bool
	}{
		{
			name:        "inside bounds",
			x:           2,
			y:           3,
			shouldWrite: true,
		},
		{
			name:        "top-left corner",
			x:           0,
			y:           0,
			shouldWrite: true,
		},
		{
			name:        "bottom-right corner",
			x:           9,
			y:           9,
			shouldWrite: true,
		},
		{
			name:        "negative X",
			x:           -1,
			y:           5,
			shouldWrite: false,
		},
		{
			name:        "negative Y",
			x:           5,
			y:           -1,
			shouldWrite: false,
		},
		{
			name:        "out of bounds X",
			x:           10,
			y:           5,
			shouldWrite: false,
		},
		{
			name:        "out of bounds Y",
			x:           5,
			y:           10,
			shouldWrite: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf.Clear()
			canvas.SetCell(tt.x, tt.y, testCell)

			if tt.shouldWrite {
				assert.Equal(t, testCell, buf.Data[tt.y][tt.x])
			} else {
				// Убеждаемся, что ячейка не изменилась (осталась пробелом)
				if tt.x >= 0 && tt.x < buf.W && tt.y >= 0 && tt.y < buf.H {
					assert.Equal(t, NewCell(' '), buf.Data[tt.y][tt.x])
				}
			}
		})
	}
}

func TestCanvas_SubCanvas(t *testing.T) {
	buf, err := NewBuffer(30, 20)
	require.NoError(t, err)
	parent := NewCanvas(buf)

	tests := []struct {
		name           string
		x, y, w, h     int
		wantOffsetX    int
		wantOffsetY    int
		wantClipW      int
		wantClipH      int
		drawLocalX     int
		drawLocalY     int
		wantGlobalX    int
		wantGlobalY    int
		shouldDrawCell bool
	}{
		{
			name:           "child fully inside parent",
			x:              5,
			y:              4,
			w:              10,
			h:              8,
			wantOffsetX:    5,
			wantOffsetY:    4,
			wantClipW:      10,
			wantClipH:      8,
			drawLocalX:     1,
			drawLocalY:     2,
			wantGlobalX:    6,
			wantGlobalY:    6,
			shouldDrawCell: true,
		},
		{
			name:           "child exceeds parent right edge",
			x:              25,
			y:              5,
			w:              10,
			h:              5,
			wantOffsetX:    25,
			wantOffsetY:    5,
			wantClipW:      5, // 30 - 25 = 5
			wantClipH:      5,
			drawLocalX:     4,
			drawLocalY:     0,
			wantGlobalX:    29,
			wantGlobalY:    5,
			shouldDrawCell: true,
		},
		{
			name:           "child drawing beyond clipped edge",
			x:              25,
			y:              5,
			w:              10,
			h:              5,
			wantOffsetX:    25,
			wantOffsetY:    5,
			wantClipW:      5,
			wantClipH:      5,
			drawLocalX:     7, // 7 >= clipW (5) -> must be clipped!
			drawLocalY:     0,
			shouldDrawCell: false,
		},
		{
			name:        "child outside parent completely",
			x:           35,
			y:           25,
			w:           10,
			h:           10,
			wantOffsetX: 35,
			wantOffsetY: 25,
			wantClipW:   0,
			wantClipH:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf.Clear()
			sub := parent.SubCanvas(tt.x, tt.y, tt.w, tt.h)

			assert.Equal(t, tt.wantOffsetX, sub.offsetX)
			assert.Equal(t, tt.wantOffsetY, sub.offsetY)
			assert.Equal(t, tt.wantClipW, sub.clipW)
			if tt.wantClipH >= 0 {
				assert.Equal(t, tt.wantClipH, sub.clipH)
			}

			if tt.shouldDrawCell {
				ColorRed, _ := BaseColors.Get("red")
				cell := NewCellColored('M', ColorRed, ColorDefault())
				sub.SetCell(tt.drawLocalX, tt.drawLocalY, cell)
				assert.Equal(t, cell, buf.Data[tt.wantGlobalY][tt.wantGlobalX])
			}
		})
	}
}

func TestCanvas_SetRune(t *testing.T) {
	buf, err := NewBuffer(5, 5)
	require.NoError(t, err)
	canvas := NewCanvas(buf)

	tests := []struct {
		name string
		x, y int
		r    rune
	}{
		{name: "set custom rune", x: 1, y: 1, r: 'Q'},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf.Clear()
			canvas.SetRune(tt.x, tt.y, tt.r)
			assert.Equal(t, NewCell(tt.r), buf.Data[tt.y][tt.x])
		})
	}
}

func TestCanvas_DrawString(t *testing.T) {
	tests := []struct {
		name       string
		bufW, bufH int
		setup      func(c *Canvas)
		verify     func(t *testing.T, b *Buffer)
	}{
		{
			name: "renders string horizontally",
			bufW: 10, bufH: 5,
			setup: func(c *Canvas) {
				c.DrawString(2, 1, "GO", ColorDefault())
			},
			verify: func(t *testing.T, b *Buffer) {
				assert.Equal(t, 'G', b.Data[1][2].R)
				assert.Equal(t, 'O', b.Data[1][3].R)
				assert.Equal(t, ' ', b.Data[1][4].R)
			},
		},
		{
			name: "clips string at canvas boundary",
			bufW: 10, bufH: 5,
			setup: func(c *Canvas) {
				sub := c.SubCanvas(0, 0, 4, 3)
				sub.DrawString(0, 0, "TOOLONG", ColorDefault())
			},
			verify: func(t *testing.T, b *Buffer) {
				assert.Equal(t, 'T', b.Data[0][0].R)
				assert.Equal(t, 'O', b.Data[0][1].R)
				assert.Equal(t, 'O', b.Data[0][2].R)
				assert.Equal(t, 'L', b.Data[0][3].R)
				assert.Equal(t, ' ', b.Data[0][4].R) // 5-й символ обрезан
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf, err := NewBuffer(tt.bufW, tt.bufH)
			require.NoError(t, err)
			canvas := NewCanvas(buf)

			tt.setup(canvas)
			tt.verify(t, buf)
		})
	}
}

func TestCanvas_DrawRect(t *testing.T) {
	tests := []struct {
		name       string
		x, y, w, h int
		style      BorderType
		verify     func(t *testing.T, b *Buffer)
	}{
		{
			name: "ignores dimensions smaller than 2",
			x:    0, y: 0, w: 1, h: 5,
			style: BorderSingle,
			verify: func(t *testing.T, b *Buffer) {
				assert.Equal(t, NewCell(' '), b.Data[0][0])
			},
		},
		{
			name: "draws single border correctly",
			x:    1, y: 1, w: 4, h: 3,
			style: BorderSingle,
			verify: func(t *testing.T, b *Buffer) {
				// Углы
				assert.Equal(t, '┌', b.Data[1][1].R)
				assert.Equal(t, '┐', b.Data[1][4].R)
				assert.Equal(t, '└', b.Data[3][1].R)
				assert.Equal(t, '┘', b.Data[3][4].R)
				// Линии
				assert.Equal(t, '─', b.Data[1][2].R)
				assert.Equal(t, '│', b.Data[2][1].R)
			},
		},
		{
			name: "draws rounded border",
			x:    0, y: 0, w: 3, h: 3,
			style: BorderRounded,
			verify: func(t *testing.T, b *Buffer) {
				assert.Equal(t, '╭', b.Data[0][0].R)
				assert.Equal(t, '╯', b.Data[2][2].R)
			},
		},
		{
			name: "draws double border",
			x:    0, y: 0, w: 3, h: 3,
			style: BorderDouble,
			verify: func(t *testing.T, b *Buffer) {
				assert.Equal(t, '╔', b.Data[0][0].R)
				assert.Equal(t, '╝', b.Data[2][2].R)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf, err := NewBuffer(10, 10)
			require.NoError(t, err)
			canvas := NewCanvas(buf)

			canvas.DrawRect(tt.x, tt.y, tt.w, tt.h, tt.style, ColorDefault(), ColorDefault())
			tt.verify(t, buf)
		})
	}
}
