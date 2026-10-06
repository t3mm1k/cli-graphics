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

		ta.renderMarkdownLine(canvas, textX, textY, textW, wrappedLines[lineIndex], style)
	}
}

// renderMarkdownLine разбирает строку на наличие `#` (заголовки) и `**` (жирный текст)
// и посимвольно выводит её на холст с соответствующими атрибутами стилей.
func (ta *TextArea) renderMarkdownLine(canvas *engine.Canvas, startX, startY, maxW int, line string, baseStyle engine.Style) {
	runes := []rune(line)
	if len(runes) == 0 {
		return
	}

	isHeader := false
	if len(runes) >= 2 && runes[0] == '#' && runes[1] == ' ' {
		isHeader = true
		runes = runes[2:]
	}

	currentStyle := baseStyle.TextStyle
	currentFg := baseStyle.Fg

	if isHeader {
		currentStyle = engine.TextStyleBold
		// // можно подсветить заголовок желтым
		// if yellowColor, ok := engine.BaseColors.Get("yellow"); ok {
		// 	currentFg = yellowColor
		// }
	}

	x := startX
	actualLen := len(runes)
	
	switch baseStyle.Align {
	case engine.TextAlignCenter:
		x = startX + (maxW-actualLen)/2
	case engine.TextAlignRight:
		x = startX + maxW - actualLen
	}

	isBoldMode := false
	visualIdx := 0

	for j := 0; j < len(runes); j++ {
		if visualIdx >= maxW {
			break
		}

		if !isHeader && j < len(runes)-1 && runes[j] == '*' && runes[j+1] == '*' {
			isBoldMode = !isBoldMode
			j++
			continue
		}

		finalStyle := currentStyle
		if isBoldMode {
			finalStyle = engine.TextStyleBold
		}

		cell := engine.NewCellColored(runes[j], currentFg, baseStyle.Bg)
		cell.Style = finalStyle

		canvas.SetCell(x+visualIdx, startY, cell)
		visualIdx++
	}
}

func (ta *TextArea) OnTick() {}
