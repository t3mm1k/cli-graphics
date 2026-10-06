package widgets

import (
	"cli-graphics/engine"
	"strings"

	"github.com/google/uuid"
)

type TextArea struct {
	engine.BaseComponent

	rawLines []string
	
	topLine int
	focused bool
}

func NewTextArea(w, h, x, y int, initialText string, classes ...string) *TextArea {
	id := uuid.New()
	rawLines := strings.Split(initialText, "\n")

	ta := &TextArea{
		BaseComponent: engine.NewBaseComponent(id, x, y, w, h),
		rawLines:     rawLines,
		topLine:      0,
		focused:      false,
	}

	ta.InitStyle("border-single border:dim-gray fg:white bg:default focus:border:neon-pink", classes...)

	engine.Registry.AddComponent(ta)
	engine.FocusManagerInstance.Register(id)

	return ta
}

// wrapText принимает доступную ширину и нарезает исходный текст на строки,
// аккуратно перенося слова по пробелам, чтобы они не разрывались пополам
func (ta *TextArea) wrapText(maxLineWidth int) []string {
	if maxLineWidth <= 0 {
		return nil
	}

	var wrappedLines []string

	for _, rawLine := range ta.rawLines {
		if len(rawLine) == 0 {
			wrappedLines = append(wrappedLines, "")
			continue
		}

		words := strings.Fields(rawLine)
		if len(words) == 0 {
			wrappedLines = append(wrappedLines, "")
			continue
		}

		var currentLine strings.Builder

		for _, word := range words {
			if len(word) > maxLineWidth {
				if currentLine.Len() > 0 {
					wrappedLines = append(wrappedLines, currentLine.String())
					currentLine.Reset()
				}
				runes := []rune(word)
				for len(runes) > maxLineWidth {
					wrappedLines = append(wrappedLines, string(runes[:maxLineWidth]))
					runes = runes[maxLineWidth:]
				}
				currentLine.WriteString(string(runes))
				continue
			}

			spaceNeeded := 0
			if currentLine.Len() > 0 {
				spaceNeeded = 1
			}

			if currentLine.Len()+spaceNeeded+len(word) > maxLineWidth {
				wrappedLines = append(wrappedLines, currentLine.String())
				currentLine.Reset()
				currentLine.WriteString(word)
			} else {
				if spaceNeeded > 0 {
					currentLine.WriteString(" ")
				}
				currentLine.WriteString(word)
			}
		}

		if currentLine.Len() > 0 {
			wrappedLines = append(wrappedLines, currentLine.String())
		}
	}

	return wrappedLines
}

func (ta *TextArea) SetFocus(focused bool) {
	ta.focused = focused
}

func (ta *TextArea) IsFocused() bool {
	return engine.FocusManagerInstance.GetFocused() == ta.GetId()
}

func (ta *TextArea) HandleKey(key string) bool {
	style := ta.CurrentStyle()
	w, h := ta.GetSize()
	
	offset := 1
	if style.Border == engine.BorderNone {
		offset = 0
	}
	
	textH := h - (offset * 2) - style.Padding.Top - style.Padding.Bottom
	wrappedLines := ta.wrapText(w - (offset * 2) - style.Padding.Left - style.Padding.Right)
	
	if len(wrappedLines) <= textH {
		return false
	}

	switch key {
	case "Down", "ArrowDown":
		if ta.topLine < len(wrappedLines)-textH {
			ta.topLine++
			return true
		}
	case "Up", "ArrowUp":
		if ta.topLine > 0 {
			ta.topLine--
			return true
		}
	}
	return false
}

func (ta *TextArea) Render(canvas *engine.Canvas) {
	w, h := ta.GetSize()
	style := ta.CurrentStyle()

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

	wrappedLines := ta.wrapText(textW)

	for i := 0; i < textH; i++ {
		lineIndex := ta.topLine + i
		if lineIndex >= len(wrappedLines) {
			break
		}

		textX := offset + style.Padding.Left
		switch style.Align {
		case engine.TextAlignCenter:
			textX = offset + style.Padding.Left + (textW)/2
		case engine.TextAlignRight:
			textX = w - offset - style.Padding.Right
		}
		textY := offset + style.Padding.Top + i

		canvas.DrawStringAligned(textX, textY, wrappedLines[lineIndex], style.Fg, style.TextStyle, style.Align)
	}
}

func (ta *TextArea) OnTick() {}
