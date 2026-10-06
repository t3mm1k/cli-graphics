package widgets

import (
	"cli-graphics/engine"

	"unicode/utf8"
)

type InputArea struct {
	TextArea

	cursorX         int
	cursorY         int
	cursorIsVisible bool
}

func NewInputArea(w, h, x, y int, initialText string, classes ...string) *InputArea {
	baseTA := NewTextArea(w, h, x, y, initialText, classes...)
	
	ia := &InputArea{
		TextArea:        *baseTA,
		cursorX:         0,
		cursorY:         0,
		cursorIsVisible: true,
	}

	ia.InitStyle("border-single border:dim-gray fg:white bg:default focus:border:neon-green", classes...)

	engine.Registry.AddComponent(ia)
	engine.FocusManagerInstance.Register(ia.GetId())

	return ia
}

func (ia *InputArea) HandleKey(key string) bool {
	style := ia.CurrentStyle()
	w, h := ia.GetSize()

	offset := 1
	if style.Border == engine.BorderNone {
		offset = 0
	}

	textW := w - (offset * 2) - style.Padding.Left - style.Padding.Right
	textH := h - (offset * 2) - style.Padding.Top - style.Padding.Bottom

	switch key {
	case "Up", "ArrowUp":
		if ia.cursorY > 0 {
			ia.cursorY--
			lineRunes := []rune(ia.rawLines[ia.cursorY])
			if ia.cursorX > len(lineRunes) {
				ia.cursorX = len(lineRunes)
			}
			ia.ensureCursorVisible(textH)
			return true
		}
	case "Down", "ArrowDown":
		if ia.cursorY < len(ia.rawLines)-1 {
			ia.cursorY++
			lineRunes := []rune(ia.rawLines[ia.cursorY])
			if ia.cursorX > len(lineRunes) {
				ia.cursorX = len(lineRunes)
			}
			ia.ensureCursorVisible(textH)
			return true
		}
	case "Left", "ArrowLeft":
		if ia.cursorX > 0 {
			ia.cursorX--
			return true
		} else if ia.cursorY > 0 {
			ia.cursorY--
			ia.cursorX = len([]rune(ia.rawLines[ia.cursorY]))
			ia.ensureCursorVisible(textH)
			return true
		}
	case "Right", "ArrowRight":
		lineRunes := []rune(ia.rawLines[ia.cursorY])
		if ia.cursorX < len(lineRunes) {
			ia.cursorX++
			return true
		} else if ia.cursorY < len(ia.rawLines)-1 {
			ia.cursorY++
			ia.cursorX = 0
			ia.ensureCursorVisible(textH)
			return true
		}
	case "Enter":
		lineRunes := []rune(ia.rawLines[ia.cursorY])
		leftPart := string(lineRunes[:ia.cursorX])
		rightPart := string(lineRunes[ia.cursorX:])

		ia.rawLines[ia.cursorY] = leftPart
		ia.rawLines = append(ia.rawLines[:ia.cursorY+1], append([]string{rightPart}, ia.rawLines[ia.cursorY+1:]...)...)
		
		ia.cursorY++
		ia.cursorX = 0
		ia.ensureCursorVisible(textH)
		return true

	case "Backspace":
		lineRunes := []rune(ia.rawLines[ia.cursorY])
		if ia.cursorX > 0 {
			newRunes := append(lineRunes[:ia.cursorX-1], lineRunes[ia.cursorX:]...)
			ia.rawLines[ia.cursorY] = string(newRunes)
			ia.cursorX--
			return true
		} else if ia.cursorY > 0 {
			prevLineRunes := []rune(ia.rawLines[ia.cursorY-1])
			oldCursorX := len(prevLineRunes)

			ia.rawLines[ia.cursorY-1] = ia.rawLines[ia.cursorY-1] + ia.rawLines[ia.cursorY]
			ia.rawLines = append(ia.rawLines[:ia.cursorY], ia.rawLines[ia.cursorY+1:]...)

			ia.cursorY--
			ia.cursorX = oldCursorX
			ia.ensureCursorVisible(textH)
			return true
		}
	default:
		if utf8.RuneCountInString(key) == 1 {
			char, _ := utf8.DecodeRuneInString(key)
			lineRunes := []rune(ia.rawLines[ia.cursorY])
			
			if len(lineRunes) >= textW {
				leftPart := string(lineRunes[:ia.cursorX])
				rightPart := string(lineRunes[ia.cursorX:])

				ia.rawLines[ia.cursorY] = leftPart
				ia.rawLines = append(ia.rawLines[:ia.cursorY+1], append([]string{rightPart}, ia.rawLines[ia.cursorY+1:]...)...)
				
				ia.cursorY++
				ia.cursorX = 0
				ia.ensureCursorVisible(textH)
				
				lineRunes = []rune(ia.rawLines[ia.cursorY])
			}

			newRunes := append(lineRunes[:ia.cursorX], append([]rune{char}, lineRunes[ia.cursorX:]...)...)
			ia.rawLines[ia.cursorY] = string(newRunes)
			ia.cursorX++
			return true
		}
	}

	return false
}

func (ia *InputArea) ensureCursorVisible(textH int) {
	if ia.cursorY < ia.topLine {
		ia.topLine = ia.cursorY
	} else if ia.cursorY >= ia.topLine+textH {
		ia.topLine = ia.cursorY - textH + 1
	}
}

func (ia *InputArea) OnTick() {
	if ia.IsFocused() {
		ia.cursorIsVisible = !ia.cursorIsVisible
	}
}

func (ia *InputArea) Render(canvas *engine.Canvas) {
	w, h := ia.GetSize()
	style := ia.CurrentStyle()

	canvas.DrawRect(0, 0, w, h, style.Border, style.BorderFg, style.Bg)

	offset := 1
	if style.Border == engine.BorderNone {
		offset = 0
	}

	textW := w - (offset * 2) - style.Padding.Left - style.Padding.Right
	textH := h - (offset * 2) - style.Padding.Top - style.Padding.Bottom

	if textW <= 0 || textH <= 0 {
		return
	}

	for i := 0; i < textH; i++ {
		lineIndex := ia.topLine + i
		if lineIndex >= len(ia.rawLines) {
			break
		}

		textX := offset + style.Padding.Left
		textY := offset + style.Padding.Top + i

		runes := []rune(ia.rawLines[lineIndex])
		
		if len(runes) > textW {
			runes = runes[:textW]
		}

		canvas.DrawStringAligned(textX, textY, string(runes), style.Fg, style.TextStyle, style.Align)
	}

	if ia.IsFocused() && ia.cursorIsVisible {
		if ia.cursorY >= ia.topLine && ia.cursorY < ia.topLine+textH {
			visualY := offset + style.Padding.Top + (ia.cursorY - ia.topLine)
			visualX := offset + style.Padding.Left + ia.cursorX
			
			if visualX < w-offset-style.Padding.Right {
				cursorColor := style.BorderFg
				if cursorColor.IsDefault {
					cursorColor = engine.ColorDefault()
				}
				cursorCell := engine.NewCellColored('_', cursorColor, style.Bg)
				canvas.SetCell(visualX, visualY, cursorCell)
			}
		}
	}
}
