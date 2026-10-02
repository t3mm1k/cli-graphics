package engine

import (
	"strconv"
	"strings"
)

type BorderType int

const (
	BorderNone BorderType = iota
	BorderSingle
	BorderRounded
	BorderDouble
	BorderBold
)

type TextAlign int

const (
	TextAlignLeft TextAlign = iota
	TextAlignCenter
	TextAlignRight
)

type TextOverflow int

const (
	OverflowClip TextOverflow = iota
	OverflowTruncate
)

type Padding struct {
	Top    int
	Right  int
	Bottom int
	Left   int
}

type Style struct {
	Fg        Color
	Bg        Color
	TextStyle TextStyle

	Border   BorderType
	BorderFg Color
	BorderBg Color

	Padding  Padding
	Align    TextAlign
	Overflow TextOverflow
}

func DefaultStyle() Style {
	return Style{
		Fg:        ColorDefault(),
		Bg:        ColorDefault(),
		TextStyle: TextStyleDefault,
		Border:    BorderNone,
		BorderFg:  ColorDefault(),
		BorderBg:  ColorDefault(),
		Padding:   Padding{},
		Align:     TextAlignLeft,
		Overflow:  OverflowClip,
	}
}

func ParseColor(val string) (Color, bool) {
	val = strings.ToLower(strings.TrimSpace(val))

	if val == "default" || val == "none" {
		return ColorDefault(), true
	}

	if color, ok := BaseColors.Get(val); ok {
		return color, true
	}

	if strings.HasPrefix(val, "#") {
		hexStr := strings.TrimPrefix(val, "#")
		if len(hexStr) == 6 {
			rgb, err := strconv.ParseUint(hexStr, 16, 32)
			if err == nil {
				return Color{
					R:         uint8((rgb >> 16) & 0xFF),
					G:         uint8((rgb >> 8) & 0xFF),
					B:         uint8(rgb & 0xFF),
					IsDefault: false,
				}, true
			}
		}
	}

	return ColorDefault(), false
}
