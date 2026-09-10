package main

import (
	"cli-graphics/widgets"
)

func main() {
	rootScreen := widgets.NewBox(0, 0, 80, 24)

	window1 := widgets.NewBox(5, 5, 30, 10)
	window2 := widgets.NewBox(40, 5, 20, 8)
	list := widgets.NewList(12, 5, 29, 2, []string{"тп на аме — Серега Пират", "Почему ты еще не фанат? — Серега Пират", "Я поднимаю свою голову вверх — Серега Пират", "ЧСВ — Lida & Серега Пират", "Зомби апокалипсис — Серега Пират", "Вайбмен — Серега Пират", "как же он силён — Серега Пират", "Ну и что, что я вор? — Серега Пират", "прости я не знаю — Серега Пират", "ну где моя нога — Серега Пират"})
	label := widgets.NewLabel(40, 10, "jopa", 2)

	rootScreen.AddChild(window1)
	rootScreen.AddChild(window2)
	rootScreen.AddChild(list)
	rootScreen.AddChild(label)

	rootScreen.Render()

	rootScreen.Buffer.Flush()

}
