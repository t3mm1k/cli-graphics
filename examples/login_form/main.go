package main

import (
	"fmt"

	"cli-graphics/engine"
	"cli-graphics/widgets"
)

func main() {
	// Фоновый контейнер на весь экран
	root := widgets.NewBox(0, 0, 80, 24, "border-single border:dark-gray bg:default")

	// Диалоговое окно авторизации по центру
	dialog := widgets.NewBox(18, 3, 44, 18, "border-double border:neon-cyan bg:default")
	root.AddChild(dialog)

	// Заголовок диалога
	title := widgets.NewLabel(24, 4, false, "🔐  USER AUTHENTICATION", "fg:yellow bold")
	subtitle := widgets.NewLabel(22, 5, false, "Please enter your credentials below", "fg:dim-gray italic")
	root.AddChild(title)
	root.AddChild(subtitle)

	// Поле ввода логина
	usernameLabel := widgets.NewLabel(21, 7, false, "Username / Email:", "fg:white bold")
	root.AddChild(usernameLabel)

	usernameInput := widgets.NewInput(21, 8, 38, func(key string) {},
		"border-single border:dim-gray fg:white focus:border:neon-cyan")
	root.AddChild(usernameInput)

	// Поле ввода пароля (со скрытыми символами '*')
	passwordLabel := widgets.NewLabel(21, 11, false, "Password:", "fg:white bold")
	root.AddChild(passwordLabel)

	passwordInput := widgets.NewInput(21, 12, 38, func(key string) {},
		"border-single border:dim-gray fg:white focus:border:neon-cyan")
	passwordInput.SetPassword(true)
	root.AddChild(passwordInput)

	// Лейбл статуса / ошибок
	statusLabel := widgets.NewLabel(21, 15, false, "Ready to log in", "fg:dim-gray")
	root.AddChild(statusLabel)

	// Кнопка входа
	btnLogin := widgets.NewButton(21, 17, "  Login  ", func() {
		username := string(usernameInput.GetValue())
		password := string(passwordInput.GetValue())

		if len(username) == 0 {
			statusLabel.SetText("❌ Error: Username cannot be empty!")
			statusLabel.InitStyle("fg:red bold")
			return
		}
		if len(password) < 4 {
			statusLabel.SetText("❌ Error: Password must be >= 4 chars!")
			statusLabel.InitStyle("fg:red bold")
			return
		}

		statusLabel.SetText(fmt.Sprintf("✔ Welcome back, %s!", username))
		statusLabel.InitStyle("fg:pastel-green bold")
	}, "border-rounded border:dim-gray fg:white focus:border:yellow active:reverse")

	// Кнопка очистки
	btnClear := widgets.NewButton(34, 17, "  Clear  ", func() {
		usernameInput.SetValue([]rune{})
		passwordInput.SetValue([]rune{})
		statusLabel.SetText("Fields cleared")
		statusLabel.InitStyle("fg:dim-gray")
	}, "border-rounded border:dim-gray fg:coral focus:border:red active:reverse")

	root.AddChild(btnLogin)
	root.AddChild(btnClear)

	// Подсказка внизу экрана
	hint := widgets.NewLabel(2, 23, false, "[Tab] Next field  •  [Shift+Tab] Prev field  •  [Enter] Click  •  [Ctrl+C] Exit", "fg:dark-gray")
	root.AddChild(hint)

	app := engine.NewApp(root)
	if err := app.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
