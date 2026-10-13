package widgets

import (
	"fmt"
	"testing"
	"time"

	"github.com/t3mm1k/cli-graphics/engine"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProgressBar_Initialization(t *testing.T) {
	bar := NewProgressBar(0, 0, 20, 3)
	require.NotNil(t, bar)

	assert.Equal(t, 0.0, bar.GetProgress())
	assert.Equal(t, FormatPercentage, bar.GetFormat())
	assert.True(t, bar.IsShowText())

	w, h := bar.GetSize()
	assert.Equal(t, 20, w)
	assert.Equal(t, 3, h)
}

func TestProgressBar_SetProgressClamping(t *testing.T) {
	bar := NewProgressBar(0, 0, 20, 3)

	bar.SetProgress(0.5)
	assert.Equal(t, 0.5, bar.GetProgress())

	bar.SetProgress(-0.2)
	assert.Equal(t, 0.0, bar.GetProgress())

	bar.SetProgress(1.5)
	assert.Equal(t, 1.0, bar.GetProgress())
}

func TestProgressBar_SetTime(t *testing.T) {
	bar := NewProgressBar(0, 0, 20, 3)

	bar.SetTime(30*time.Second, 60*time.Second)
	assert.Equal(t, 0.5, bar.GetProgress())
	cur, total := bar.GetTime()
	assert.Equal(t, 30*time.Second, cur)
	assert.Equal(t, 60*time.Second, total)

	// Total 0
	bar.SetTime(10*time.Second, 0)
	assert.Equal(t, 0.0, bar.GetProgress())

	// Current exceeds total
	bar.SetTime(120*time.Second, 60*time.Second)
	assert.Equal(t, 1.0, bar.GetProgress())
}

func TestProgressBar_FormatText(t *testing.T) {
	bar := NewProgressBar(0, 0, 20, 3)
	bar.SetTime(45*time.Second, 90*time.Second) // 50%

	// 1. Percentage
	bar.SetFormat(FormatPercentage)
	assert.Equal(t, "50%", bar.FormatText())

	// 2. Time (mm:ss)
	bar.SetFormat(FormatTime)
	assert.Equal(t, "00:45 / 01:30", bar.FormatText())

	// Hours formatting
	bar.SetTime(3665*time.Second, 7200*time.Second)
	assert.Equal(t, "01:01:05 / 02:00:00", bar.FormatText())

	// 3. Seconds
	bar.SetTime(45*time.Second, 90*time.Second)
	bar.SetFormat(FormatSeconds)
	assert.Equal(t, "45s / 90s", bar.FormatText())

	// 4. Custom
	bar.SetFormat(FormatCustom)
	bar.SetCustomFormatter(func(current, total time.Duration, percent float64) string {
		return fmt.Sprintf("ETA: %.0fs", (total - current).Seconds())
	})
	assert.Equal(t, "ETA: 45s", bar.FormatText())
}

func TestProgressBar_Render(t *testing.T) {
	buf, err := engine.NewBuffer(22, 3)
	require.NoError(t, err)

	canvas := engine.NewCanvas(buf)

	bar := NewProgressBar(0, 0, 20, 3, "border-single")
	bar.SetRunes('=', '-')
	bar.SetProgress(0.5) // innerW = 18, filled = 9
	bar.SetShowText(false)

	bar.Render(canvas)

	// Check inner line (y = 1)
	// col 0 is border '│'
	// col 1..9 should be '='
	// col 10..18 should be '-'
	// col 19 is border '│'
	data := buf.GetObjects()
	assert.Equal(t, '│', data[1][0].R)
	for col := 1; col <= 9; col++ {
		assert.Equal(t, '=', data[1][col].R, "col %d should be '='", col)
	}
	for col := 10; col <= 18; col++ {
		assert.Equal(t, '-', data[1][col].R, "col %d should be '-'", col)
	}
	assert.Equal(t, '│', data[1][19].R)
}
