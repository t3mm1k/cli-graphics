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
}

func (i *Input) IsFocused() bool {
	return i.isFocused
}

func (i *Input) HandleKey(key string) bool {
	if key == "Tab" {
		return false
	} else if key == "Enter" {

	} else if key == "Backspace" {
		// TODO пофиксить бекспейс на русской раскладке(
		if len(i.value) == 0 {
			return true
		}
		i.value = i.value[:len(i.value)-1]
	} else {

		if i.OnInput != nil {
			i.OnInput(key)
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
