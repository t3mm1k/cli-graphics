package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBorderType_String(t *testing.T) {
	tests := []struct {
		name string
		b    BorderType
		want string
	}{
		{
			name: "borderType to string: none",
			b:    BorderNone,
			want: "border-none",
		},
		{
			name: "borderType to string: single",
			b:    BorderSingle,
			want: "border-single",
		},
		{
			name: "borderType to string: double",
			b:    BorderDouble,
			want: "border-double",
		},
		{
			name: "borderType to string: rounded",
			b:    BorderRounded,
			want: "border-rounded",
		},
		{
			name: "borderType to string: bold",
			b:    BorderBold,
			want: "border-bold",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.b.String(), "String()")
		})
	}
}

func TestDefaultStyle(t *testing.T) {
	tests := []struct {
		name string
		want Style
	}{
		{
			name: "defaultStyle",
			want: Style{
				Fg:        ColorDefault(),
				Bg:        ColorDefault(),
				TextStyle: TextStyleDefault,
				Border:    BorderNone,
				BorderFg:  ColorDefault(),
				BorderBg:  ColorDefault(),
				Padding:   Padding{},
				Align:     TextAlignLeft,
				Overflow:  OverflowClip,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, DefaultStyle(), "DefaultStyle()")
		})
	}
}

func TestParseColor(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name  string
		args  args
		want  Color
		want1 bool
	}{
		{
			name: "ParseColor from baseColors",
			args: args{
				val: "charcoal",
			},
			want:  Color{R: 26, G: 26, B: 26, IsDefault: false},
			want1: true,
		},
		{
			name: "ParseColor from HEX",
			args: args{
				val: "#00FFFF",
			},
			want:  Color{R: 0, G: 255, B: 255, IsDefault: false},
			want1: true,
		},
		{
			name: "ParseColor default",
			args: args{
				val: "default",
			},
			want:  ColorDefault(),
			want1: true,
		},
		{
			name: "ParseColor invalid value",
			args: args{
				val: "invalid",
			},
			want:  ColorDefault(),
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ParseColor(tt.args.val)
			assert.Equalf(t, tt.want, got, "ParseColor(%v)", tt.args.val)
			assert.Equalf(t, tt.want1, got1, "ParseColor(%v)", tt.args.val)
		})
	}
}

func TestParseStyles(t *testing.T) {
	type args struct {
		base    Style
		classes []string
	}
	tests := []struct {
		name         string
		args         args
		wantNormal   Style
		wantFocus    Style
		wantDisabled Style
		wantActive   Style
		wantSelected Style
		wantHover    Style
	}{
		{
			name: "empty classes returns base for all states",
			args: args{
				base:    DefaultStyle(),
				classes: []string{},
			},
			wantNormal:   DefaultStyle(),
			wantFocus:    DefaultStyle(),
			wantDisabled: DefaultStyle(),
			wantActive:   DefaultStyle(),
			wantSelected: DefaultStyle(),
			wantHover:    DefaultStyle(),
		},
		{
			name: "basic classes applied and inherited by all states",
			args: args{
				base:    DefaultStyle(),
				classes: []string{"fg:yellow border-rounded px-2"},
			},
			wantNormal: Style{
				Fg:        Color{R: 255, G: 255, B: 0, IsDefault: false},
				Bg:        ColorDefault(),
				TextStyle: TextStyleDefault,
				Border:    BorderRounded,
				BorderFg:  ColorDefault(),
				BorderBg:  ColorDefault(),
				Padding:   Padding{Left: 2, Right: 2},
				Align:     TextAlignLeft,
				Overflow:  OverflowClip,
			},
			wantFocus: Style{
				Fg:        Color{R: 255, G: 255, B: 0, IsDefault: false},
				Bg:        ColorDefault(),
				TextStyle: TextStyleDefault,
				Border:    BorderRounded,
				BorderFg:  ColorDefault(),
				BorderBg:  ColorDefault(),
				Padding:   Padding{Left: 2, Right: 2},
				Align:     TextAlignLeft,
				Overflow:  OverflowClip,
			},
			wantDisabled: Style{
				Fg:        Color{R: 255, G: 255, B: 0, IsDefault: false},
				Bg:        ColorDefault(),
				TextStyle: TextStyleDefault,
				Border:    BorderRounded,
				BorderFg:  ColorDefault(),
				BorderBg:  ColorDefault(),
				Padding:   Padding{Left: 2, Right: 2},
				Align:     TextAlignLeft,
				Overflow:  OverflowClip,
			},
			wantActive: Style{
				Fg:        Color{R: 255, G: 255, B: 0, IsDefault: false},
				Bg:        ColorDefault(),
				TextStyle: TextStyleDefault,
				Border:    BorderRounded,
				BorderFg:  ColorDefault(),
				BorderBg:  ColorDefault(),
				Padding:   Padding{Left: 2, Right: 2},
				Align:     TextAlignLeft,
				Overflow:  OverflowClip,
			},
			wantSelected: Style{
				Fg:        Color{R: 255, G: 255, B: 0, IsDefault: false},
				Bg:        ColorDefault(),
				TextStyle: TextStyleDefault,
				Border:    BorderRounded,
				BorderFg:  ColorDefault(),
				BorderBg:  ColorDefault(),
				Padding:   Padding{Left: 2, Right: 2},
				Align:     TextAlignLeft,
				Overflow:  OverflowClip,
			},
			wantHover: Style{
				Fg:        Color{R: 255, G: 255, B: 0, IsDefault: false},
				Bg:        ColorDefault(),
				TextStyle: TextStyleDefault,
				Border:    BorderRounded,
				BorderFg:  ColorDefault(),
				BorderBg:  ColorDefault(),
				Padding:   Padding{Left: 2, Right: 2},
				Align:     TextAlignLeft,
				Overflow:  OverflowClip,
			},
		},
		{
			name: "state overrides: focus, disabled, active, selected, hover",
			args: args{
				base:    DefaultStyle(),
				classes: []string{"fg:white border-single focus:border:yellow focus:bold disabled:dim active:bg:blue selected:underline hover:border:red hover:italic"},
			},
			wantNormal: Style{
				Fg:        Color{R: 255, G: 255, B: 255, IsDefault: false},
				Bg:        ColorDefault(),
				TextStyle: TextStyleDefault,
				Border:    BorderSingle,
				BorderFg:  ColorDefault(),
				BorderBg:  ColorDefault(),
				Padding:   Padding{},
				Align:     TextAlignLeft,
				Overflow:  OverflowClip,
			},
			wantFocus: Style{
				Fg:        Color{R: 255, G: 255, B: 255, IsDefault: false},
				Bg:        ColorDefault(),
				TextStyle: TextStyleBold,
				Border:    BorderSingle,
				BorderFg:  Color{R: 255, G: 255, B: 0, IsDefault: false},
				BorderBg:  ColorDefault(),
				Padding:   Padding{},
				Align:     TextAlignLeft,
				Overflow:  OverflowClip,
			},
			wantDisabled: Style{
				Fg:        Color{R: 255, G: 255, B: 255, IsDefault: false},
				Bg:        ColorDefault(),
				TextStyle: TextStyleDim,
				Border:    BorderSingle,
				BorderFg:  ColorDefault(),
				BorderBg:  ColorDefault(),
				Padding:   Padding{},
				Align:     TextAlignLeft,
				Overflow:  OverflowClip,
			},
			wantActive: Style{
				Fg:        Color{R: 255, G: 255, B: 255, IsDefault: false},
				Bg:        Color{R: 0, G: 120, B: 255, IsDefault: false},
				TextStyle: TextStyleDefault,
				Border:    BorderSingle,
				BorderFg:  ColorDefault(),
				BorderBg:  ColorDefault(),
				Padding:   Padding{},
				Align:     TextAlignLeft,
				Overflow:  OverflowClip,
			},
			wantSelected: Style{
				Fg:        Color{R: 255, G: 255, B: 255, IsDefault: false},
				Bg:        ColorDefault(),
				TextStyle: TextStyleUnderline,
				Border:    BorderSingle,
				BorderFg:  ColorDefault(),
				BorderBg:  ColorDefault(),
				Padding:   Padding{},
				Align:     TextAlignLeft,
				Overflow:  OverflowClip,
			},
			wantHover: Style{
				Fg:        Color{R: 255, G: 255, B: 255, IsDefault: false},
				Bg:        ColorDefault(),
				TextStyle: TextStyleItalic,
				Border:    BorderSingle,
				BorderFg:  Color{R: 255, G: 0, B: 0, IsDefault: false},
				BorderBg:  ColorDefault(),
				Padding:   Padding{},
				Align:     TextAlignLeft,
				Overflow:  OverflowClip,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNormal, gotFocus, gotDisabled, gotActive, gotSelected, gotHover := ParseStyles(tt.args.base, tt.args.classes)
			assert.Equalf(t, tt.wantNormal, gotNormal, "ParseStyles(%v, %v)", tt.args.base, tt.args.classes)
			assert.Equalf(t, tt.wantFocus, gotFocus, "ParseStyles(%v, %v)", tt.args.base, tt.args.classes)
			assert.Equalf(t, tt.wantDisabled, gotDisabled, "ParseStyles(%v, %v)", tt.args.base, tt.args.classes)
			assert.Equalf(t, tt.wantActive, gotActive, "ParseStyles(%v, %v)", tt.args.base, tt.args.classes)
			assert.Equalf(t, tt.wantSelected, gotSelected, "ParseStyles(%v, %v)", tt.args.base, tt.args.classes)
			assert.Equalf(t, tt.wantHover, gotHover, "ParseStyles(%v, %v)", tt.args.base, tt.args.classes)
		})
	}
}

func TestTextAlign_String(t *testing.T) {
	tests := []struct {
		name string
		b    TextAlign
		want string
	}{
		{
			name: "textAlign to string: left",
			b:    TextAlignLeft,
			want: "text-left",
		},
		{
			name: "textAlign to string: right",
			b:    TextAlignRight,
			want: "text-right",
		},
		{
			name: "textAlign to string: center",
			b:    TextAlignCenter,
			want: "text-center",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.b.String(), "String()")
		})
	}
}

func TestTextOverflow_String(t *testing.T) {
	tests := []struct {
		name string
		b    TextOverflow
		want string
	}{
		{
			name: "textOverflow to string: clip",
			b:    OverflowClip,
			want: "clip",
		},
		{
			name: "textOverflow to string: truncate",
			b:    OverflowTruncate,
			want: "truncate",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.b.String(), "String()")
		})
	}
}

func TestTextStyle_String(t *testing.T) {
	tests := []struct {
		name string
		b    TextStyle
		want string
	}{
		{
			name: "textStyle to string: bold",
			b:    TextStyleBold,
			want: "bold",
		},
		{
			name: "textStyle to string: italic",
			b:    TextStyleItalic,
			want: "italic",
		},
		{
			name: "textStyle to string: underline",
			b:    TextStyleUnderline,
			want: "underline",
		},
		{
			name: "textStyle to string: blink",
			b:    TextStyleBlink,
			want: "blink",
		},
		{
			name: "textStyle to string: dim",
			b:    TextStyleDim,
			want: "dim",
		},
		{
			name: "textStyle to string: reverse",
			b:    TextStyleReverse,
			want: "reverse",
		},
		{
			name: "textStyle to string: hidden",
			b:    TextStyleHidden,
			want: "hidden",
		},
		{
			name: "textStyle to string: strikethrough",
			b:    TextStyleStrikethrough,
			want: "strikethrough",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.b.String(), "String()")
		})
	}
}

func Test_applyRule(t *testing.T) {
	type args struct {
		s    *Style
		rule string
	}
	tests := []struct {
		name string
		args args
		want Style
	}{
		{
			name: "border type rule",
			args: args{s: &Style{}, rule: "border-rounded"},
			want: Style{Border: BorderRounded},
		},
		{
			name: "align rule",
			args: args{s: &Style{}, rule: "text-center"},
			want: Style{Align: TextAlignCenter},
		},
		{
			name: "overflow rule",
			args: args{s: &Style{}, rule: "truncate"},
			want: Style{Overflow: OverflowTruncate},
		},
		{
			name: "textstyle rule",
			args: args{s: &Style{}, rule: "bold"},
			want: Style{TextStyle: TextStyleBold},
		},
		{
			name: "padding rule",
			args: args{s: &Style{}, rule: "px-2"},
			want: Style{Padding: Padding{Left: 2, Right: 2}},
		},
		{
			name: "fg color rule",
			args: args{s: &Style{}, rule: "fg:yellow"},
			want: Style{Fg: Color{R: 255, G: 255, B: 0, IsDefault: false}},
		},
		{
			name: "bg color rule",
			args: args{s: &Style{}, rule: "bg:black"},
			want: Style{Bg: Color{R: 0, G: 0, B: 0, IsDefault: false}},
		},
		{
			name: "border fg rule",
			args: args{s: &Style{}, rule: "border:red"},
			want: Style{BorderFg: Color{R: 255, G: 0, B: 0, IsDefault: false}},
		},
		{
			name: "border bg rule",
			args: args{s: &Style{}, rule: "border-bg:white"},
			want: Style{BorderBg: Color{R: 255, G: 255, B: 255, IsDefault: false}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			applyRule(tt.args.s, tt.args.rule)
			assert.Equalf(t, tt.want, *tt.args.s, "applyRule(%v)", tt.args.rule)
		})
	}
}

func Test_parseColorRule(t *testing.T) {
	type args struct {
		s   *Style
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "fg:yellow", args: args{s: &Style{}, str: "fg:yellow"}, want: true},
		{name: "bg:black", args: args{s: &Style{}, str: "bg:black"}, want: true},
		{name: "border:red", args: args{s: &Style{}, str: "border:red"}, want: true},
		{name: "border-bg:white", args: args{s: &Style{}, str: "border-bg:white"}, want: true},
		{name: "hex fg", args: args{s: &Style{}, str: "fg:#123456"}, want: true},
		{name: "invalid color", args: args{s: &Style{}, str: "fg:notacolor"}, want: false},
		{name: "invalid prefix", args: args{s: &Style{}, str: "color:red"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, parseColorRule(tt.args.s, tt.args.str), "parseColorRule(%v, %v)", tt.args.s, tt.args.str)
		})
	}
}

func Test_parsePadding(t *testing.T) {
	type args struct {
		s   *Style
		str string
	}
	tests := []struct {
		name        string
		args        args
		want        bool
		wantPadding Padding
	}{
		{name: "p-2", args: args{s: &Style{}, str: "p-2"}, want: true, wantPadding: Padding{Top: 2, Right: 2, Bottom: 2, Left: 2}},
		{name: "px-3", args: args{s: &Style{}, str: "px-3"}, want: true, wantPadding: Padding{Left: 3, Right: 3}},
		{name: "py-4", args: args{s: &Style{}, str: "py-4"}, want: true, wantPadding: Padding{Top: 4, Bottom: 4}},
		{name: "pt-1", args: args{s: &Style{}, str: "pt-1"}, want: true, wantPadding: Padding{Top: 1}},
		{name: "pb-5", args: args{s: &Style{}, str: "pb-5"}, want: true, wantPadding: Padding{Bottom: 5}},
		{name: "pl-6", args: args{s: &Style{}, str: "pl-6"}, want: true, wantPadding: Padding{Left: 6}},
		{name: "pr-7", args: args{s: &Style{}, str: "pr-7"}, want: true, wantPadding: Padding{Right: 7}},
		{name: "invalid format", args: args{s: &Style{}, str: "p-abc"}, want: false, wantPadding: Padding{}},
		{name: "unknown prefix", args: args{s: &Style{}, str: "margin-2"}, want: false, wantPadding: Padding{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parsePadding(tt.args.s, tt.args.str)
			assert.Equalf(t, tt.want, got, "parsePadding(%v, %v)", tt.args.s, tt.args.str)
			if tt.want {
				assert.Equalf(t, tt.wantPadding, tt.args.s.Padding, "Padding mismatch")
			}
		})
	}
}

func Test_str2align(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name  string
		args  args
		want  TextAlign
		want1 bool
	}{
		{name: "left", args: args{str: "text-left"}, want: TextAlignLeft, want1: true},
		{name: "center", args: args{str: "text-center"}, want: TextAlignCenter, want1: true},
		{name: "right", args: args{str: "text-right"}, want: TextAlignRight, want1: true},
		{name: "invalid", args: args{str: "text-justify"}, want: TextAlignCenter, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := str2align(tt.args.str)
			assert.Equalf(t, tt.want, got, "str2align(%v)", tt.args.str)
			assert.Equalf(t, tt.want1, got1, "str2align(%v)", tt.args.str)
		})
	}
}

func Test_str2bordertype(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name  string
		args  args
		want  BorderType
		want1 bool
	}{
		{name: "single", args: args{str: "border-single"}, want: BorderSingle, want1: true},
		{name: "rounded", args: args{str: "border-rounded"}, want: BorderRounded, want1: true},
		{name: "double", args: args{str: "border-double"}, want: BorderDouble, want1: true},
		{name: "bold", args: args{str: "border-bold"}, want: BorderBold, want1: true},
		{name: "none", args: args{str: "border-none"}, want: BorderNone, want1: true},
		{name: "invalid", args: args{str: "border-dotted"}, want: BorderNone, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := str2bordertype(tt.args.str)
			assert.Equalf(t, tt.want, got, "str2bordertype(%v)", tt.args.str)
			assert.Equalf(t, tt.want1, got1, "str2bordertype(%v)", tt.args.str)
		})
	}
}

func Test_str2overflow(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name  string
		args  args
		want  TextOverflow
		want1 bool
	}{
		{name: "truncate", args: args{str: "truncate"}, want: OverflowTruncate, want1: true},
		{name: "clip", args: args{str: "clip"}, want: OverflowClip, want1: true},
		{name: "invalid", args: args{str: "ellipsis"}, want: OverflowClip, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := str2overflow(tt.args.str)
			assert.Equalf(t, tt.want, got, "str2overflow(%v)", tt.args.str)
			assert.Equalf(t, tt.want1, got1, "str2overflow(%v)", tt.args.str)
		})
	}
}

func Test_str2textstyle(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name  string
		args  args
		want  TextStyle
		want1 bool
	}{
		{name: "bold", args: args{str: "bold"}, want: TextStyleBold, want1: true},
		{name: "dim", args: args{str: "dim"}, want: TextStyleDim, want1: true},
		{name: "italic", args: args{str: "italic"}, want: TextStyleItalic, want1: true},
		{name: "underline", args: args{str: "underline"}, want: TextStyleUnderline, want1: true},
		{name: "blink", args: args{str: "blink"}, want: TextStyleBlink, want1: true},
		{name: "reverse", args: args{str: "reverse"}, want: TextStyleReverse, want1: true},
		{name: "hidden", args: args{str: "hidden"}, want: TextStyleHidden, want1: true},
		{name: "strikethrough", args: args{str: "strikethrough"}, want: TextStyleStrikethrough, want1: true},
		{name: "invalid", args: args{str: "shadow"}, want: TextStyleDefault, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := str2textstyle(tt.args.str)
			assert.Equalf(t, tt.want, got, "str2textstyle(%v)", tt.args.str)
			assert.Equalf(t, tt.want1, got1, "str2textstyle(%v)", tt.args.str)
		})
	}
}
