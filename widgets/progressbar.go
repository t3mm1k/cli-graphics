package widgets

import (
	"fmt"
	"time"

	"cli-graphics/engine"

	"github.com/google/uuid"
)

type ProgressFormat int

const (
	FormatPercentage ProgressFormat = iota // "50%"
	FormatTime                             // "01:15 / 02:30"
	FormatSeconds                          // "45s / 120s"
	FormatCustom                           // User Format
)

const DefaultProgressBarClasses = "border-single border:dim-gray fg:default bg:default text-center"

type ProgressBar struct {
	engine.BaseComponent

	progress float64
	current  time.Duration
	total    time.Duration

	format    ProgressFormat
	showText  bool
	fillRune  rune
	emptyRune rune

	customFormatter func(current, total time.Duration, percent float64) string
}

func NewProgressBar(x, y, w, h int, classes ...string) *ProgressBar {
	id := uuid.New()
	bar := &ProgressBar{
		BaseComponent: engine.NewBaseComponent(id, x, y, w, h),
		progress:      0.0,
		format:        FormatPercentage,
		showText:      true,
		fillRune:      '=',
		emptyRune:     '-',
	}

	bar.InitStyle(DefaultProgressBarClasses, classes...)
	engine.Registry.AddComponent(bar)

	return bar
}

func (p *ProgressBar) SetProgress(progress float64) {
	if progress < 0.0 {
		progress = 0.0
	} else if progress > 1.0 {
		progress = 1.0
	}
	p.progress = progress
}

func (p *ProgressBar) GetProgress() float64 {
	return p.progress
}

func (p *ProgressBar) SetTime(current, total time.Duration) {
	if total < 0 {
		total = 0
	}
	if current < 0 {
		current = 0
	} else if total > 0 && current > total {
		current = total
	}
	p.current = current
	p.total = total
	if total > 0 {
		p.progress = float64(current) / float64(total)
	} else {
		p.progress = 0.0
	}
}

func (p *ProgressBar) GetTime() (current, total time.Duration) {
	return p.current, p.total
}

func (p *ProgressBar) SetFormat(format ProgressFormat) {
	p.format = format
}

func (p *ProgressBar) GetFormat() ProgressFormat {
	return p.format
}

func (p *ProgressBar) SetShowText(show bool) {
	p.showText = show
}

func (p *ProgressBar) IsShowText() bool {
	return p.showText
}

func (p *ProgressBar) SetRunes(fill, empty rune) {
	p.fillRune = fill
	p.emptyRune = empty
}

func (p *ProgressBar) SetCustomFormatter(fn func(current, total time.Duration, percent float64) string) {
	p.customFormatter = fn
}

func formatDuration(d time.Duration) string {
	seconds := int(d.Seconds())
	h := seconds / 3600
	m := (seconds % 3600) / 60
	s := seconds % 60
	if h > 0 {
		return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
	}
	return fmt.Sprintf("%02d:%02d", m, s)
}

func (p *ProgressBar) FormatText() string {
	switch p.format {
	case FormatPercentage:
		return fmt.Sprintf("%d%%", int(p.progress*100))
	case FormatTime:
		return fmt.Sprintf("%s / %s", formatDuration(p.current), formatDuration(p.total))
	case FormatSeconds:
		return fmt.Sprintf("%ds / %ds", int(p.current.Seconds()), int(p.total.Seconds()))
	case FormatCustom:
		if p.customFormatter != nil {
			return p.customFormatter(p.current, p.total, p.progress)
		}
		return fmt.Sprintf("%d%%", int(p.progress*100))
	default:
		return fmt.Sprintf("%d%%", int(p.progress*100))
	}
}

func (p *ProgressBar) OnTick() {}

func (p *ProgressBar) Render(canvas *engine.Canvas) {
	style := p.CurrentStyle()
	w, h := p.GetSize()
	if w <= 0 || h <= 0 {
		return
	}

	offset := 0
	if style.Border != engine.BorderNone {
		canvas.DrawRect(0, 0, w, h, style.Border, style.BorderFg, style.Bg)
		offset = 1
	} else if !style.Bg.IsDefault {
		canvas.Fill(0, 0, w, h, engine.NewCellColored(' ', engine.ColorDefault(), style.Bg))
	}

	innerW := w - 2*offset
	innerH := h - 2*offset
	if innerW <= 0 || innerH <= 0 {
		return
	}

	filledCols := int(float64(innerW) * p.progress)
	if filledCols < 0 {
		filledCols = 0
	} else if filledCols > innerW {
		filledCols = innerW
	}

	emptyColor := style.BorderFg
	if emptyColor.IsDefault {
		if c, ok := engine.BaseColors.Get("dim-gray"); ok {
			emptyColor = c
		} else {
			emptyColor = engine.ColorDefault()
		}
	}

	// Отрисовка полосы прогресса (заполненная и незаполненная части)
	for row := 0; row < innerH; row++ {
		y := offset + row
		for col := 0; col < innerW; col++ {
			x := offset + col
			if col < filledCols {
				cell := engine.NewCellColored(p.fillRune, style.Fg, style.Bg)
				canvas.SetCell(x, y, cell)
			} else {
				cell := engine.NewCellColored(p.emptyRune, emptyColor, style.Bg)
				canvas.SetCell(x, y, cell)
			}
		}
	}

	// Наложение текста поверх полосы (если включено)
	if p.showText {
		text := p.FormatText()
		if text != "" {
			textY := offset + innerH/2
			textX := offset + innerW/2
			switch style.Align {
			case engine.TextAlignLeft:
				textX = offset + 1
			case engine.TextAlignRight:
				textX = offset + innerW - 1
			case engine.TextAlignCenter:
				textX = offset + innerW/2
			}

			textColor := style.Fg
			if textColor.IsDefault {
				if c, ok := engine.BaseColors.Get("white"); ok {
					textColor = c
				} else {
					textColor = engine.ColorDefault()
				}
			}
			textStyle := style.TextStyle
			if textStyle == engine.TextStyleDefault {
				textStyle = engine.TextStyleBold
			}

			canvas.DrawStringAligned(textX, textY, text, textColor, textStyle, style.Align)
		}
	}
}
