package utils

import (
	"bytes"
	"strconv"
)

type MouseAction int

const (
	MouseActionPress MouseAction = iota
	MouseActionRelease
	MouseActionMove
)

type MouseBtn int

const (
	MouseBtnLeft MouseBtn = iota
	MouseBtnMiddle
	MouseBtnRight
	MouseBtnWheelUp
	MouseBtnWheelDown
	MouseBtnNone
)

type SGRMouseData struct {
	Button MouseBtn
	Action MouseAction
	X, Y   int
}

// ParseSGRMouse attempts to parse an ANSI SGR 1006 mouse sequence: "\x1b[<btn;x;y(M|m)"
func ParseSGRMouse(data []byte) (*SGRMouseData, bool) {
	if !bytes.HasPrefix(data, []byte("\x1b[<")) {
		return nil, false
	}
	if len(data) < 7 {
		return nil, false
	}

	finalChar := data[len(data)-1]
	if finalChar != 'M' && finalChar != 'm' {
		return nil, false
	}

	content := string(data[3 : len(data)-1])
	parts := bytes.Split([]byte(content), []byte(";"))
	if len(parts) != 3 {
		return nil, false
	}

	rawBtn, err := strconv.Atoi(string(parts[0]))
	if err != nil {
		return nil, false
	}
	x, err := strconv.Atoi(string(parts[1]))
	if err != nil {
		return nil, false
	}
	y, err := strconv.Atoi(string(parts[2]))
	if err != nil {
		return nil, false
	}

	if x > 0 {
		x -= 1
	}
	if y > 0 {
		y -= 1
	}

	var action MouseAction
	var btn MouseBtn

	isMotion := (rawBtn & 32) != 0
	btnCode := rawBtn &^ 32

	if finalChar == 'm' {
		action = MouseActionRelease
	} else if isMotion {
		action = MouseActionMove
	} else {
		action = MouseActionPress
	}

	switch btnCode {
	case 0:
		btn = MouseBtnLeft
	case 1:
		btn = MouseBtnMiddle
	case 2:
		btn = MouseBtnRight
	case 64:
		btn = MouseBtnWheelUp
	case 65:
		btn = MouseBtnWheelDown
	default:
		btn = MouseBtnNone
	}

	return &SGRMouseData{
		Button: btn,
		Action: action,
		X:      x,
		Y:      y,
	}, true
}
