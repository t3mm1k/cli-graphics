package widgets

import (
	"cli-graphics/engine"
	"unicode"

	"github.com/google/uuid"
)

const DefaultInputClasses = "border-single border:dim-gray fg:white focus:border-double disabled:dim hover:border:white"

type Input struct {
	engine.BaseComponent

	value []rune

	cursorIsVisible bool

	cursorPos int

	OnInput func(key string)

	Filter func(key string) bool

	isPassword bool
}

func (i *Input) SetFocus(focused bool) {
	if focused {
		i.cursorIsVisible = true
	} else {
		i.cursorIsVisible = false
	}
}

func (i *Input) GetValue() []rune {
	return i.value
}

func (i *Input) SetValue(value []rune) {
	i.value = value
	if i.cursorPos > len(value) {
		i.cursorPos = len(value)
	}
}

func (i *Input) IsFocused() bool {
	return engine.FocusManagerInstance.GetFocused() == i.GetId()
}

func (i *Input) HandleKey(key string) bool {
	if key == "Tab" {
		return false
	}

	switch key {
	case "Enter":

	case "Backspace":
		if len(i.value) == 0 || i.cursorPos == 0 {
			return true
		}
		i.value = append(i.value[:i.cursorPos-1], i.value[i.cursorPos:]...)
		i.cursorPos--

	case "Left", "ArrowLeft":
		if i.cursorPos > 0 {
			i.cursorPos--
		}
		return true

	case "Right", "ArrowRight":
		if i.cursorPos < len(i.value) {
			i.cursorPos++
		}
		return true

	default:
		runesKey := []rune(key)
		if len(runesKey) != 1 {
			return true
		}

		if i.Filter != nil && !i.Filter(key) {
			return true
		}

		i.value = append(i.value[:i.cursorPos], append([]rune{runesKey[0]}, i.value[i.cursorPos:]...)...)
		i.cursorPos++

		if i.OnInput != nil {
			i.OnInput(key)
		}
	}

	return true
}

func (i *Input) OnTick() {
	if i.IsFocused() {
		i.cursorIsVisible = !i.cursorIsVisible
	}
}

func (i *Input) Render(canvas *engine.Canvas) {
	w, h := i.GetSize()
	style := i.CurrentStyle()

	canvas.DrawRect(0, 0, w, h, style.Border, style.BorderFg, style.Bg)

	offset := 1
	if style.Border == engine.BorderNone {
		offset = 0
	}

	textW := w - (offset * 2) - style.Padding.Left - style.Padding.Right
	if textW <= 0 {
		return
	}

	startX := offset + style.Padding.Left
	startY := offset + style.Padding.Top

	for col := 0; col < textW; col++ {
		canvas.SetCell(startX+col, startY, engine.NewCellColored(' ', style.Fg, style.Bg))
	}

	start := 0
	if i.cursorPos >= textW {
		start = i.cursorPos - textW + 1
	}

	for j := 0; j < textW; j++ {
		strIndex := start + j
		if strIndex < len(i.value) {
			charToRender := i.value[strIndex]
			if i.isPassword {
				charToRender = '*'
			}

			cell := engine.NewCellColored(charToRender, style.Fg, style.Bg)
			cell.Style = style.TextStyle
			canvas.SetCell(startX+j, startY, cell)
		}
	}

	if i.IsFocused() && i.cursorIsVisible {
		visualCursorPos := i.cursorPos - start
		if visualCursorPos >= 0 && visualCursorPos < textW {
			cursorColor := style.BorderFg
			if cursorColor.IsDefault {
				cursorColor = engine.ColorDefault()
			}
			cursorCell := engine.NewCellColored('_', cursorColor, style.Bg)
			canvas.SetCell(startX+visualCursorPos, startY, cursorCell)
		}
	}
}

func NewInput(x, y, w int, onInput func(key string), classes ...string) *Input {
	id := uuid.New()
	input := &Input{
		BaseComponent:   engine.NewBaseComponent(id, x, y, w, 3),
		OnInput:         onInput,
		cursorIsVisible: false,
		cursorPos:       0,
		Filter: func(key string) bool {
			runes := []rune(key)
			if len(runes) != 1 {
				return false
			}
			r := runes[0]
			return unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsPunct(r) || unicode.IsSymbol(r) || r == ' '
		},
	}

	input.InitStyle(DefaultInputClasses, classes...)

	engine.Registry.AddComponent(input)
	engine.FocusManagerInstance.Register(id)

	return input
}

func (i *Input) SetPassword(isPassword bool) {
	i.isPassword = isPassword
}

func (i *Input) SetColors(normal, focus engine.Color) {
	i.Style.BorderFg = normal
	i.FocusedStyle.BorderFg = focus
}
