package main

import (
	"cli-graphics/ui"
	"cli-graphics/ui/listView"
	"cli-graphics/ui/window"
)

func main() {
	win := window.New(17, 61)
	list := listView.New(12, 30, ui.Point{X: 29, Y: 2}, []string{"тп на аме — Серега Пират", "Почему ты еще не фанат? — Серега Пират", "Я поднимаю свою голову вверх — Серега Пират", "ЧСВ — Lida & Серега Пират", "Зомби апокалипсис — Серега Пират", "Вайбмен — Серега Пират", "как же он силён — Серега Пират", "Ну и что, что я вор? — Серега Пират", "прости я не знаю — Серега Пират", "ну где моя нога — Серега Пират"})
	win.AddObject((*listView.ListView)(list))
	win.Render()

	win.Print()
}
