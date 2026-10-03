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

func str2bordertype(str string) (BorderType, bool) {
	switch str {
	case "border-single":
		return BorderSingle, true
	case "border-rounded":
		return BorderRounded, true
	case "border-double":
		return BorderDouble, true
	case "border-bold":
		return BorderBold, true
	case "border-none":
		return BorderNone, true
	default:
		return BorderNone, false
	}
}
func (b BorderType) String() string {
	switch b {
	case BorderSingle:
		return "border-single"
	case BorderRounded:
		return "border-rounded"
	case BorderDouble:
		return "border-double"
	case BorderBold:
		return "border-bold"
	case BorderNone:
		return "border-none"
	default:
		return ""
	}
}

type TextAlign int

const (
	TextAlignLeft TextAlign = iota
	TextAlignCenter
	TextAlignRight
)

func str2align(str string) (TextAlign, bool) {
	switch str {
	case "text-left":
		return TextAlignLeft, true
	case "text-right":
		return TextAlignRight, true
	case "text-center":
		return TextAlignCenter, true
	default:
		return TextAlignCenter, false
	}
}
func (b TextAlign) String() string {
	switch b {
	case TextAlignLeft:
		return "text-left"
	case TextAlignRight:
		return "text-right"
	case TextAlignCenter:
		return "text-center"
	default:
		return ""
	}
}

type TextOverflow int

const (
	OverflowClip TextOverflow = iota
	OverflowTruncate
)

func str2overflow(str string) (TextOverflow, bool) {
	switch str {
	case "truncate":
		return OverflowTruncate, true
	case "clip":
		return OverflowClip, true
	default:
		return OverflowClip, false
	}
}
func (b TextOverflow) String() string {
	switch b {
	case OverflowTruncate:
		return "truncate"
	default:
		return "clip"
	}
}

type TextStyle uint8

const (
	TextStyleDefault TextStyle = iota
	TextStyleBold
	TextStyleDim
	TextStyleItalic
	TextStyleUnderline
	TextStyleBlink
	TextStyleReverse
	TextStyleHidden
	TextStyleStrikethrough
)

func str2textstyle(str string) (TextStyle, bool) {
	switch str {
	case "bold":
		return TextStyleBold, true
	case "dim":
		return TextStyleDim, true
	case "italic":
		return TextStyleItalic, true
	case "underline":
		return TextStyleUnderline, true
	case "blink":
		return TextStyleBlink, true
	case "reverse":
		return TextStyleReverse, true
	case "hidden":
		return TextStyleHidden, true
	case "strikethrough":
		return TextStyleStrikethrough, true
	default:
		return TextStyleDefault, false
	}
}
func (b TextStyle) String() string {
	switch b {
	case TextStyleBold:
		return "bold"
	case TextStyleDim:
		return "dim"
	case TextStyleItalic:
		return "italic"
	case TextStyleUnderline:
		return "underline"
	case TextStyleBlink:
		return "blink"
	case TextStyleReverse:
		return "reverse"
	case TextStyleHidden:
		return "hidden"
	case TextStyleStrikethrough:
		return "strikethrough"
	default:
		return ""
	}
}

type Padding struct {
	Top    int
	Right  int
	Bottom int
	Left   int
}

func parsePadding(s *Style, str string) bool {

	var prefix string
	switch {
	case strings.HasPrefix(str, "px-"):
		prefix = "px-"
	case strings.HasPrefix(str, "py-"):
		prefix = "py-"
	case strings.HasPrefix(str, "pt-"):
		prefix = "pt-"
	case strings.HasPrefix(str, "pb-"):
		prefix = "pb-"
	case strings.HasPrefix(str, "pl-"):
		prefix = "pl-"
	case strings.HasPrefix(str, "pr-"):
		prefix = "pr-"
	case strings.HasPrefix(str, "p-"):
		prefix = "p-"
	default:
		return false
	}

	valStr := strings.TrimPrefix(str, prefix)
	val, err := strconv.Atoi(valStr)
	if err != nil || val < 0 {
		return false
	}

	switch prefix {
	case "p-":
		s.Padding.Top = val
		s.Padding.Right = val
		s.Padding.Bottom = val
		s.Padding.Left = val
	case "px-":
		s.Padding.Left = val
		s.Padding.Right = val
	case "py-":
		s.Padding.Top = val
		s.Padding.Bottom = val
	case "pt-":
		s.Padding.Top = val
	case "pb-":
		s.Padding.Bottom = val
	case "pl-":
		s.Padding.Left = val
	case "pr-":
		s.Padding.Right = val
	}

	return true

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

func parseColorRule(s *Style, str string) bool {
	var prefix string
	switch {
	case strings.HasPrefix(str, "bg:"):
		prefix = "bg:"
	case strings.HasPrefix(str, "fg:"):
		prefix = "fg:"
	case strings.HasPrefix(str, "border:"):
		prefix = "border:"
	case strings.HasPrefix(str, "border-bg:"):
		prefix = "border-bg:"
	default:
		return false
	}

	val := strings.TrimPrefix(str, prefix)
	if color, ok := ParseColor(val); ok {
		switch prefix {
		case "bg:":
			s.Bg = color
		case "fg:":
			s.Fg = color
		case "border:":
			s.BorderFg = color
		case "border-bg:":
			s.BorderBg = color
		default:
			return false
		}
		return true
	}
	return false
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

	if after, ok := strings.CutPrefix(val, "#"); ok {
		hexStr := after
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

func applyRule(s *Style, rule string) {
	if border, ok := str2bordertype(rule); ok {
		s.Border = border
	}
	if align, ok := str2align(rule); ok {
		s.Align = align
	}
	if overflow, ok := str2overflow(rule); ok {
		s.Overflow = overflow
	}
	if textStyle, ok := str2textstyle(rule); ok {
		s.TextStyle = textStyle
	}

	parsePadding(s, rule)
	parseColorRule(s, rule)
	// TODO Переписать функции, привести к одному виду
}

func ParseStyles(base Style, classes []string) (normal, focus, disabled, active, selected, hover Style) {
	normal = base
	classStr := strings.Join(classes, " ")
	tokens := strings.Fields(classStr)
	for _, class := range tokens {
		if !strings.Contains(class, ":") || strings.HasPrefix(class, "fg:") || strings.HasPrefix(class, "bg:") || strings.HasPrefix(class, "border:") || strings.HasPrefix(class, "border-bg:") {
			applyRule(&normal, class)
		}
	}
	focus = normal
	disabled = normal
	active = normal
	selected = normal
	hover = normal
	for _, class := range tokens {
		switch {
		case strings.HasPrefix(class, "focus:"):
			applyRule(&focus, strings.TrimPrefix(class, "focus:"))
		case strings.HasPrefix(class, "disabled:"):
			applyRule(&disabled, strings.TrimPrefix(class, "disabled:"))
		case strings.HasPrefix(class, "active:"):
			applyRule(&active, strings.TrimPrefix(class, "active:"))
		case strings.HasPrefix(class, "selected:"):
			applyRule(&selected, strings.TrimPrefix(class, "selected:"))
		case strings.HasPrefix(class, "hover:"):
			applyRule(&hover, strings.TrimPrefix(class, "hover:"))
		}
	}
	return normal, focus, disabled, active, selected, hover
}
