package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseKey(t *testing.T) {
	tests := []struct {
		name string
		raw  []byte
		want string
	}{
		{
			name: "empty bytes",
			raw:  []byte{},
			want: "",
		},
		{
			name: "single char ASCII",
			raw:  []byte("a"),
			want: "a",
		},
		{
			name: "space",
			raw:  []byte(" "),
			want: " ",
		},
		{
			name: "tab",
			raw:  []byte{9},
			want: "Tab",
		},
		{
			name: "enter line feed",
			raw:  []byte{10},
			want: "Enter",
		},
		{
			name: "enter carriage return",
			raw:  []byte{13},
			want: "Enter",
		},
		{
			name: "escape",
			raw:  []byte{27},
			want: "Escape",
		},
		{
			name: "backspace 8",
			raw:  []byte{8},
			want: "Backspace",
		},
		{
			name: "backspace 127",
			raw:  []byte{127},
			want: "Backspace",
		},
		{
			name: "ctrl+c",
			raw:  []byte{3},
			want: "Ctrl+C",
		},
		{
			name: "arrow up",
			raw:  []byte{27, '[', 'A'},
			want: "Up",
		},
		{
			name: "arrow down",
			raw:  []byte{27, '[', 'B'},
			want: "Down",
		},
		{
			name: "arrow right",
			raw:  []byte{27, '[', 'C'},
			want: "Right",
		},
		{
			name: "arrow left",
			raw:  []byte{27, '[', 'D'},
			want: "Left",
		},
		{
			name: "shift+tab",
			raw:  []byte{27, '[', 'Z'},
			want: "Shift+Tab",
		},
		{
			name: "home",
			raw:  []byte{27, '[', 'H'},
			want: "Home",
		},
		{
			name: "end",
			raw:  []byte{27, '[', 'F'},
			want: "End",
		},
		{
			name: "insert key",
			raw:  []byte{27, '[', '2', '~'},
			want: "Insert",
		},
		{
			name: "delete key",
			raw:  []byte{27, '[', '3', '~'},
			want: "Delete",
		},
		{
			name: "page up",
			raw:  []byte{27, '[', '5', '~'},
			want: "PageUp",
		},
		{
			name: "page down",
			raw:  []byte{27, '[', '6', '~'},
			want: "PageDown",
		},
		{
			name: "utf8 russian rune",
			raw:  []byte("я"),
			want: "я",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseKey(tt.raw)
			assert.Equal(t, tt.want, got)
		})
	}
}
