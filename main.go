package main

import (
	"cli-graphics/widgets"
	"os"

	"golang.org/x/term"
)

func ReadKey() string {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}

	defer term.Restore(int(os.Stdin.Fd()), oldState)

	b := make([]byte, 3)
	n, _ := os.Stdin.Read(b)

	if n == 1 {
		switch b[0] {
		case 3:
			return "Ctrl+C"
		case 13:
			return "Enter"
		case 27:
			return "Escape"
		case 127, 8:
			return "Backspace"
		default:
			return string(b[0])
		}
	} else if n == 3 && b[0] == 27 && b[1] == 91 {

		switch b[2] {
		case 65:
			return "Up"
		case 66:
			return "Down"
		case 67:
			return "Right"
		case 68:
			return "Left"
		}
	}

	return "Unknown"
}
func main() {
	keyEvents := make(chan string)

	go func() {
		for {
			key := ReadKey()
			keyEvents <- key
		}
	}()

	rootScreen := widgets.NewBox(0, 0, 80, 24)

	window1 := widgets.NewBox(5, 5, 30, 10)
	window2 := widgets.NewBox(40, 5, 20, 8)
	list := widgets.NewList(12, 5, 29, 2, []string{"тп на аме — Серега Пират", "Почему ты еще не фанат? — Серега Пират", "Я поднимаю свою голову вверх — Серега Пират", "ЧСВ — Lida & Серега Пират", "Зомби апокалипсис — Серега Пират", "Вайбмен — Серега Пират", "как же он силён — Серега Пират", "Ну и что, что я вор? — Серега Пират", "прости я не знаю — Серега Пират", "ну где моя нога — Серега Пират"})
	label := widgets.NewLabel(40, 20, true, "test with border", 12)
	labelWithoutBorder := widgets.NewLabel(40, 18, false, "test without border", 12)

	rootScreen.AddChild(window1)
	rootScreen.AddChild(window2)
	rootScreen.AddChild(list)
	rootScreen.AddChild(label)
	rootScreen.AddChild(labelWithoutBorder)

	for {
		select {
		case key := <-keyEvents:
			switch key {
			case "q":
				return
			}
		}

		rootScreen.Render()

		rootScreen.Buffer.Flush()
	}

}
