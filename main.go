package main

import (
	"cli-graphics/engine"
	"cli-graphics/widgets"
	"fmt"
)

func main() {
	rootScreen := widgets.NewBox(0, 0, 80, 24)
	window1 := widgets.NewBox(5, 6, 30, 10)
	input := widgets.NewInput(1, 2, 18, nil)
	btnExit := widgets.NewButton(40, 20, "Exit", nil)
	rootScreen.AddChild(input)
	rootScreen.AddChild(window1)
	rootScreen.AddChild(btnExit)
	rootScreen.SetFocus(true)

	app := engine.NewApp(rootScreen)

	btnExit.OnClick = func() {
		app.Stop()
	}

	if err := app.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
