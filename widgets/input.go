package widgets

import (
	"cli-graphics/engine"

	"github.com/google/uuid"
)

type Input struct {
	engine.BaseComponent

	value []rune //TODO ПОМЕНЯТЬ НА []rune

	isFocused bool

	cursorIsVisible bool

	cursorPos int

	OnInput func(key string)

	Filter  func(key string) bool
}

func (i *Input) Rerender() {
	//TODO implement me
	panic("implement me")
}

func (i *Input) SetFocus(focused bool) {
	if focused == false {
		i.cursorIsVisible = false
	}
	i.isFocused = focused
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
	return i.isFocused
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
	case "Right", "ArrowRight":
		if i.cursorPos < len(i.value) {
			i.cursorPos++
		}

	default:
		if i.Filter != nil && !i.Filter(key) {
			return true
		}

		runes := []rune(key)
		if len(runes) == 1 {

			charToInsert := runes[0]
			
			res := make([]rune, 0, len(i.value)+1)
			res = append(res, i.value[:i.cursorPos]...)
			res = append(res, charToInsert)
			res = append(res, i.value[i.cursorPos:]...)
			
			i.value = res
			i.cursorPos++

			if i.OnInput != nil {
				i.OnInput(key)
			}
		}
	}
	
	return true
}

func (i *Input) OnTick() {
	if i.isFocused {
		i.cursorIsVisible = !i.cursorIsVisible
	}
}

func (i *Input) Render() {
	i.Buffer.Clear()
	i.RenderBorder(i.IsFocused())

	w, _ := i.GetSize()
	textW := w - 2
	runes := []rune(i.value)
	start := len(runes) - textW + 1
	if start < 0 {
		start = 0
	}
	text := runes[start:]
	if i.cursorIsVisible {
		text = append(text, '_')
	}
	for j, let := range text {
		i.Buffer.Data[1][j+1] = let
	}
}

func NewInput(x, y, w int, onInput func(key string)) *Input {
	id := uuid.New()
	input := &Input{
		BaseComponent:   engine.NewBaseComponent(id, x, y, w, 3, true),
		OnInput:         onInput,
		cursorIsVisible: false,
	}

	engine.Registry.AddComponent(input)

	return input
}
