package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseSGRMouse(t *testing.T) {
	tests := []struct {
		name       string
		input      []byte
		wantValid  bool
		wantBtn    MouseBtn
		wantAction MouseAction
		wantX      int
		wantY      int
	}{
		{
			name:       "Left Click Press at 10, 20 (1-based -> 9, 19)",
			input:      []byte("\x1b[<0;10;20M"),
			wantValid:  true,
			wantBtn:    MouseBtnLeft,
			wantAction: MouseActionPress,
			wantX:      9,
			wantY:      19,
		},
		{
			name:       "Left Click Release at 10, 20",
			input:      []byte("\x1b[<0;10;20m"),
			wantValid:  true,
			wantBtn:    MouseBtnLeft,
			wantAction: MouseActionRelease,
			wantX:      9,
			wantY:      19,
		},
		{
			name:       "Right Click Press",
			input:      []byte("\x1b[<2;15;5M"),
			wantValid:  true,
			wantBtn:    MouseBtnRight,
			wantAction: MouseActionPress,
			wantX:      14,
			wantY:      4,
		},
		{
			name:       "Middle Click Press",
			input:      []byte("\x1b[<1;1;1M"),
			wantValid:  true,
			wantBtn:    MouseBtnMiddle,
			wantAction: MouseActionPress,
			wantX:      0,
			wantY:      0,
		},
		{
			name:       "Mouse Drag / Move with Left Button (btn 0 + 32)",
			input:      []byte("\x1b[<32;50;25M"),
			wantValid:  true,
			wantBtn:    MouseBtnLeft,
			wantAction: MouseActionMove,
			wantX:      49,
			wantY:      24,
		},
		{
			name:       "Wheel Up at 30, 40",
			input:      []byte("\x1b[<64;30;40M"),
			wantValid:  true,
			wantBtn:    MouseBtnWheelUp,
			wantAction: MouseActionPress,
			wantX:      29,
			wantY:      39,
		},
		{
			name:       "Wheel Down at 30, 40",
			input:      []byte("\x1b[<65;30;40M"),
			wantValid:  true,
			wantBtn:    MouseBtnWheelDown,
			wantAction: MouseActionPress,
			wantX:      29,
			wantY:      39,
		},
		{
			name:      "Invalid input - not mouse sequence",
			input:     []byte("\x1b[A"),
			wantValid: false,
		},
		{
			name:      "Malformed input - missing parts",
			input:     []byte("\x1b[<0;10M"),
			wantValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, ok := ParseSGRMouse(tt.input)
			if !tt.wantValid {
				assert.False(t, ok)
				assert.Nil(t, res)
				return
			}

			assert.True(t, ok)
			assert.NotNil(t, res)
			assert.Equal(t, tt.wantBtn, res.Button)
			assert.Equal(t, tt.wantAction, res.Action)
			assert.Equal(t, tt.wantX, res.X)
			assert.Equal(t, tt.wantY, res.Y)
		})
	}
}
