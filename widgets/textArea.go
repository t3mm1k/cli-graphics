package widgets

import (
	"cli-graphics/engine"
	"strings"

	"github.com/google/uuid"
)

type TextArea struct {
	engine.BaseComponent

	rawLines []string
}

func NewTextArea(w, h, x, y int, initialText string, classes ...string) *TextArea {
	id := uuid.New()
	
	rawLines := strings.Split(initialText, "\n")

	ta := &TextArea{
		BaseComponent: engine.NewBaseComponent(id, x, y, w, h),
		rawLines:     rawLines,
	}

	ta.InitStyle("border-single border:dim-gray fg:white bg:default", classes...)

	engine.Registry.AddComponent(ta)

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

func (ta *TextArea) Render(canvas *engine.Canvas) {}
func (ta *TextArea) OnTick() {}
