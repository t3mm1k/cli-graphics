package utils

import (
	"fmt"
	"unicode/utf8"
)

func ParseKey(b []byte) string {
	if len(b) == 0 {
		return ""
	}
	if len(b) == 1 {
		switch b[0] {
		case 3:
			return "Ctrl+C"
		case 9:
			return "Tab"
		case 10, 13:
			return "Enter"
		case 27:
			return "Escape"
		case 8, 127:
			return "Backspace"
		case 32:
			return " "
		default:
			if b[0] >= 1 && b[0] <= 26 {
				return fmt.Sprintf("Ctrl+%c", 'A'+b[0]-1)
			}
			return string(b[0])
		}
	}

	if b[0] == 27 {
		if len(b) >= 3 && b[1] == '[' {
			if len(b) == 3 {
				switch b[2] {
				case 'A':
					return "Up"
				case 'B':
					return "Down"
				case 'C':
					return "Right"
				case 'D':
					return "Left"
				case 'H':
					return "Home"
				case 'F':
					return "End"
				case 'Z':
					return "Shift+Tab"
				}
			}
			if len(b) == 4 && b[3] == '~' {
				switch b[2] {
				case '1', '7':
					return "Home"
				case '2':
					return "Insert"
				case '3':
					return "Delete"
				case '4', '8':
					return "End"
				case '5':
					return "PageUp"
				case '6':
					return "PageDown"
				}
			}
		}
		if len(b) == 3 && b[1] == 'O' {
			switch b[2] {
			case 'A':
				return "Up"
			case 'B':
				return "Down"
			case 'C':
				return "Right"
			case 'D':
				return "Left"
			case 'H':
				return "Home"
			case 'F':
				return "End"
			case 'P':
				return "F1"
			case 'Q':
				return "F2"
			case 'R':
				return "F3"
			case 'S':
				return "F4"
			}
		}
		return "Unknown"
	}
	r, size := utf8.DecodeRune(b)
	if r != utf8.RuneError && size > 0 {
		return string(r)
	}
	return "Unknown"
}
