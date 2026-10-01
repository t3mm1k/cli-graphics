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

func TestCanvas_SetRune(t *testing.T) {
	buf, err := NewBuffer(5, 5)
	require.NoError(t, err)
	canvas := NewCanvas(buf)

	canvas.SetRune(1, 1, 'Q')
	assert.Equal(t, NewCell('Q'), buf.Data[1][1])
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
				cell := NewCellColored('M', ColorRed, ColorDefault())
				sub.SetCell(tt.drawLocalX, tt.drawLocalY, cell)
				assert.Equal(t, cell, buf.Data[tt.wantGlobalY][tt.wantGlobalX])
			}
		})
	}
}

func TestCanvas_DrawString(t *testing.T) {
	buf, err := NewBuffer(10, 5)
	require.NoError(t, err)

	t.Run("renders string horizontally", func(t *testing.T) {
		buf.Clear()
		canvas := NewCanvas(buf)
		canvas.DrawString(2, 1, "GO")

		assert.Equal(t, 'G', buf.Data[1][2].R)
		assert.Equal(t, 'O', buf.Data[1][3].R)
		assert.Equal(t, ' ', buf.Data[1][4].R)
	})

	t.Run("clips string at canvas boundary", func(t *testing.T) {
		buf.Clear()
		sub := NewCanvas(buf).SubCanvas(0, 0, 4, 3)
		sub.DrawString(0, 0, "TOOLONG")

		assert.Equal(t, 'T', buf.Data[0][0].R)
		assert.Equal(t, 'O', buf.Data[0][1].R)
		assert.Equal(t, 'O', buf.Data[0][2].R)
		assert.Equal(t, 'L', buf.Data[0][3].R)
		assert.Equal(t, ' ', buf.Data[0][4].R) // 5-й символ не нарисовался
	})
}

func TestCanvas_DrawRect(t *testing.T) {
	buf, err := NewBuffer(10, 10)
	require.NoError(t, err)

	t.Run("ignores dimensions smaller than 2", func(t *testing.T) {
		buf.Clear()
		canvas := NewCanvas(buf)
		canvas.DrawRect(0, 0, 1, 5, BorderSingle, ColorDefault(), ColorDefault())
		assert.Equal(t, NewCell(' '), buf.Data[0][0])
	})

	t.Run("draws single border correctly", func(t *testing.T) {
		buf.Clear()
		canvas := NewCanvas(buf)
		canvas.DrawRect(1, 1, 4, 3, BorderSingle, ColorDefault(), ColorDefault())

		// Углы
		assert.Equal(t, '┌', buf.Data[1][1].R)
		assert.Equal(t, '┐', buf.Data[1][4].R)
		assert.Equal(t, '└', buf.Data[3][1].R)
		assert.Equal(t, '┘', buf.Data[3][4].R)

		// Горизонтальные линии
		assert.Equal(t, '─', buf.Data[1][2].R)
		assert.Equal(t, '─', buf.Data[1][3].R)
		assert.Equal(t, '─', buf.Data[3][2].R)
		assert.Equal(t, '─', buf.Data[3][3].R)

		// Вертикальные линии
		assert.Equal(t, '│', buf.Data[2][1].R)
		assert.Equal(t, '│', buf.Data[2][4].R)
	})

	t.Run("draws rounded and double borders", func(t *testing.T) {
		buf.Clear()
		canvas := NewCanvas(buf)
		canvas.DrawRect(0, 0, 3, 3, BorderRounded, ColorDefault(), ColorDefault())
		assert.Equal(t, '╭', buf.Data[0][0].R)
		assert.Equal(t, '╯', buf.Data[2][2].R)

		buf.Clear()
		canvas.DrawRect(0, 0, 3, 3, BorderDouble, ColorDefault(), ColorDefault())
		assert.Equal(t, '╔', buf.Data[0][0].R)
		assert.Equal(t, '╝', buf.Data[2][2].R)
	})
}
