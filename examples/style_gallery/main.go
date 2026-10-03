package main

import (
	"fmt"

	"cli-graphics/engine"
	"cli-graphics/widgets"
)

func main() {
	// Главное окно
	root := widgets.NewBox(0, 0, 80, 24, "border-double border:neon-cyan bg:default")

	// Заголовок
	title := widgets.NewLabel(2, 1, false, "🎨  CLI Graphics — Style & Component Gallery", "fg:yellow bold")
	subtitle := widgets.NewLabel(2, 2, false, "Demonstrating borders, colors, padding, alignment & focus states", "fg:dim-gray italic")
	root.AddChild(title)
	root.AddChild(subtitle)

	// Колонка 1: Виды рамок (Borders)
	col1Title := widgets.NewLabel(4, 4, false, "1. BORDER TYPES:", "fg:neon-cyan bold")
	root.AddChild(col1Title)

	b1 := widgets.NewLabel(4, 6, true, " Single ", "border-single border:white fg:white")
	b2 := widgets.NewLabel(16, 6, true, " Rounded ", "border-rounded border:pastel-green fg:pastel-green")
	b3 := widgets.NewLabel(4, 9, true, " Double ", "border-double border:amber fg:amber")
	b4 := widgets.NewLabel(16, 9, true, " Bold ", "border-bold border:neon-pink fg:neon-pink")
	root.AddChild(b1)
	root.AddChild(b2)
	root.AddChild(b3)
	root.AddChild(b4)

	// Колонка 2: Типографика (Typography)
	col2Title := widgets.NewLabel(32, 4, false, "2. TYPOGRAPHY & ALIGN:", "fg:neon-cyan bold")
	root.AddChild(col2Title)

	t1 := widgets.NewLabel(32, 6, false, "Bold Text", "fg:white bold")
	t2 := widgets.NewLabel(32, 7, false, "Dim / Muted Text", "fg:dim-gray dim")
	t3 := widgets.NewLabel(32, 8, false, "Italic Text", "fg:pastel-purple italic")
	t4 := widgets.NewLabel(32, 9, false, "Underlined Text", "fg:pastel-cyan underline")
	t5 := widgets.NewLabel(32, 10, false, "Reversed Colors", "reverse fg:white bg:blue")
	root.AddChild(t1)
	root.AddChild(t2)
	root.AddChild(t3)
	root.AddChild(t4)
	root.AddChild(t5)

	// Колонка 3: Цветовая палитра (Colors)
	col3Title := widgets.NewLabel(58, 4, false, "3. COLOR PALETTE:", "fg:neon-cyan bold")
	root.AddChild(col3Title)

	c1 := widgets.NewLabel(58, 6, false, "● Neon Cyan", "fg:neon-cyan")
	c2 := widgets.NewLabel(58, 7, false, "● Neon Pink", "fg:neon-pink")
	c3 := widgets.NewLabel(58, 8, false, "● Pastel Green", "fg:pastel-green")
	c4 := widgets.NewLabel(58, 9, false, "● Gold / Amber", "fg:gold")
	c5 := widgets.NewLabel(58, 10, false, "● Coral / Orange", "fg:coral")
	root.AddChild(c1)
	root.AddChild(c2)
	root.AddChild(c3)
	root.AddChild(c4)
	root.AddChild(c5)

	// Разделитель
	divider := widgets.NewLabel(2, 12, false, "────────────────────────────────────────────────────────────────────────────", "fg:dark-gray")
	root.AddChild(divider)

	// Секция интерактива: Фокусы и кнопки
	interactiveTitle := widgets.NewLabel(4, 13, false, "4. INTERACTIVE WIDGETS & FOCUS STATES (Press Tab):", "fg:neon-cyan bold")
	root.AddChild(interactiveTitle)

	counter := 0
	counterLabel := widgets.NewLabel(4, 15, false, "Click count: 0", "fg:white")
	root.AddChild(counterLabel)

	btnClick := widgets.NewButton(4, 17, "Click Me (+1)", func() {
		counter++
		counterLabel.SetText(fmt.Sprintf("Click count: %d", counter))
	}, "border-rounded border:dim-gray fg:white focus:border:yellow active:reverse")

	btnReset := widgets.NewButton(22, 17, "Reset", func() {
		counter = 0
		counterLabel.SetText("Click count: 0")
	}, "border-rounded border:dim-gray fg:coral focus:border:red active:reverse")

	root.AddChild(btnClick)
	root.AddChild(btnReset)

	inputLabel := widgets.NewLabel(40, 15, false, "Type something below:", "fg:white")
	root.AddChild(inputLabel)

	input := widgets.NewInput(40, 17, 34, func(key string) {
		// Реакция на ввод
	}, "border-single border:dim-gray fg:white focus:border:neon-cyan")
	root.AddChild(input)

	// Навигационная подсказка
	help := widgets.NewLabel(2, 22, false, "[Tab / Shift+Tab] Navigate focus  •  [Enter / Space] Activate  •  [Ctrl+C] Exit", "fg:dim-gray")
	root.AddChild(help)

	app := engine.NewApp(root)
	if err := app.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
