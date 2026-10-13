package widgets

import (
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

func (p *ProgressBar) OnTick() {}

func (p *ProgressBar) Render(canvas *engine.Canvas) {}
