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

func (p *ProgressBar) Render(canvas *engine.Canvas) {}
