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

	// Колонка 3: Интерактивный список (Selectable List)
	col3Title := widgets.NewLabel(58, 4, false, "3. INTERACTIVE LIST:", "fg:neon-cyan bold")
	root.AddChild(col3Title)

	listItems := []string{"Go / Golang", "Python", "C++ / C", "JavaScript"}
	galleryList := widgets.NewList(20, 6, 58, 6, listItems, "border-rounded border:dim-gray fg:white focus:border:neon-pink selected:fg:pastel-green")
	root.AddChild(galleryList)

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
	}, "border-rounded border:dim-gray fg:white focus:border:yellow active:reverse px-2")

	btnReset := widgets.NewButton(22, 17, "  Reset  ", func() {
		counter = 0
		counterLabel.SetText("Click count: 0")
	}, "border-rounded border:dim-gray fg:coral focus:border:red active:reverse px-2")

	root.AddChild(btnClick)
	root.AddChild(btnReset)

	inputLabel := widgets.NewLabel(40, 15, false, "Type something below:", "fg:white")
	root.AddChild(inputLabel)

	input := widgets.NewInput(40, 17, 34, func(key string) {
		// Реакция на ввод
	}, "border-single border:dim-gray fg:white focus:border:neon-cyan")
	root.AddChild(input)

	
	galleryList.OnSelect = func(index int, text string) {
		inputLabel.SetText(fmt.Sprintf("Selected language: %s", text))
	}

	// Секция многострочного текста (TextArea и InputArea)
	divider2 := widgets.NewLabel(2, 20, false, "────────────────────────────────────────────────────────────────────────────", "fg:dark-gray")
	root.AddChild(divider2)

	multiLineTitle := widgets.NewLabel(4, 21, false, "5. MULTI-LINE TEXT COMPONENTS (TextArea & InputArea):", "fg:neon-cyan bold")
	root.AddChild(multiLineTitle)

	// Текст с Markdown для TextArea
	mdText := "# Markdown Preview\n" +
		"This is a **TextArea** component.\n" +
		"It supports automatic **Word Wrap** parsing.\n" +
		"You can scroll it up and down using **Arrow keys** when focused.\n" +
		"# Another Header\n" +
		"End of document text line."

	taLabel := widgets.NewLabel(4, 23, false, "TextArea (Read-only + MD + Scroll):", "fg:dim-gray")
	root.AddChild(taLabel)

	// TextArea шириной 34 и высотой 6
	galleryTextArea := widgets.NewTextArea(34, 6, 4, 25, mdText)
	root.AddChild(galleryTextArea)

	iaLabel := widgets.NewLabel(42, 23, false, "InputArea (Multi-line editor + Scroll):", "fg:dim-gray")
	root.AddChild(iaLabel)

	// InputArea шириной 34 и высотой 6
	galleryInputArea := widgets.NewInputArea(34, 6, 42, 25, "Type text here...\nPress Enter for new line.")
	root.AddChild(galleryInputArea)

	// Навигационная подсказка снизу
	help := widgets.NewLabel(2, 34, false, "[Tab / Shift+Tab] Navigate focus  •  [Arrows] Scroll & Move Cursor  •  [Ctrl+C] Exit", "fg:dim-gray")
	root.AddChild(help)

	app := engine.NewApp(root)
	if err := app.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
