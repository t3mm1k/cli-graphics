package main

import (
	"fmt"
	"image"
	"net/http"

	_ "image/jpeg"
	_ "image/png"

	"cli-graphics/engine"
	"cli-graphics/widgets"
)

func main() {
	// 1. Главный контейнер плеера
	root := widgets.NewBox(0, 0, 80, 24, "border-double border:neon-cyan bg:default")

	// 2. Шапка плеера
	header := widgets.NewLabel(2, 1, false, "🎵  CLI Yandex Music Player  [v0.2.0]", "fg:yellow bold")
	root.AddChild(header)

	// 3. Блок с информацией о треке
	trackTitle := widgets.NewLabel(4, 3, false, "Never Gonna Give You Up", "fg:white bold")
	artistName := widgets.NewLabel(4, 4, false, "Rick Astley • Whenever You Need Somebody (1987)", "fg:dim-gray italic")
	statusLabel := widgets.NewLabel(4, 5, false, "▶ Playing  [01:23 / 03:32]", "fg:pastel-green")
	root.AddChild(trackTitle)
	root.AddChild(artistName)
	root.AddChild(statusLabel)

	// 4. Кнопки управления воспроизведением
	isPlaying := true
	btnPrev := widgets.NewButton(4, 7, "◀◀ Prev", func() {
		statusLabel.SetText("⏮ Track changed to Previous")
	}, "border-rounded fg:white border:dim-gray focus:border:yellow active:reverse")

	var btnPlay *widgets.Button
	btnPlay = widgets.NewButton(16, 7, "⏸ Pause", func() {
		isPlaying = !isPlaying
		if isPlaying {
			btnPlay.SetText("⏸ Pause")
			statusLabel.SetText("▶ Playing  [01:24 / 03:32]")
		} else {
			btnPlay.SetText("▶ Play ")
			statusLabel.SetText("⏸ Paused   [01:24 / 03:32]")
		}
	}, "border-rounded fg:yellow border:dim-gray focus:border:yellow active:reverse")

	btnNext := widgets.NewButton(28, 7, "Next ▶▶", func() {
		statusLabel.SetText("⏭ Track changed to Next")
	}, "border-rounded fg:white border:dim-gray focus:border:yellow active:reverse")

	root.AddChild(btnPrev)
	root.AddChild(btnPlay)
	root.AddChild(btnNext)

	// 5. Обложка трека (загружается асинхронно)
	coverArt := widgets.NewImage(4, 10, 32, 12, nil)
	root.AddChild(coverArt)

	// 6. Плейлист / Очередь воспроизведения справа
	playlistTitle := widgets.NewLabel(42, 3, false, "📑 Up Next in Queue:", "fg:neon-cyan bold")
	root.AddChild(playlistTitle)

	tracks := []string{
		"1. Rick Astley - Together Forever",
		"2. The Weeknd - Blinding Lights",
		"3. Daft Punk - Get Lucky",
		"4. Queen - Bohemian Rhapsody",
		"5. Michael Jackson - Billie Jean",
		"6. Depeche Mode - Enjoy the Silence",
		"7. Gorillaz - Feel Good Inc.",
	}
	playlist := widgets.NewList(35, 12, 42, 5, tracks, "border-rounded border:dim-gray fg:white focus:border:neon-pink")
	root.AddChild(playlist)

	// 7. Поле поиска треков снизу
	searchLabel := widgets.NewLabel(42, 18, false, "🔍 Search track:", "fg:dim-gray")
	root.AddChild(searchLabel)

	searchInput := widgets.NewInput(42, 19, 35, func(key string) {
		// Callback на ввод
	}, "border-single border:dim-gray fg:white focus:border:neon-green")
	root.AddChild(searchInput)

	// 8. Подсказка по навигации внизу
	helpText := widgets.NewLabel(2, 23, false, "[Tab / Shift+Tab] Switch focus  •  [Space / Enter] Click  •  [Ctrl+C] Quit", "fg:dark-gray")
	root.AddChild(helpText)

	// Запуск приложения
	app := engine.NewApp(root)

	// Загружаем тестовую обложку асинхронно
	go func() {
		imgURL := "https://avatars.yandex.net/get-music-content/16334817/73d15c0c.a.16218390-2/400x400"
		resp, err := http.Get(imgURL)
		if err != nil || resp.StatusCode != http.StatusOK {
			return
		}
		defer resp.Body.Close()

		img, _, err := image.Decode(resp.Body)
		if err != nil {
			return
		}

		app.Post(func() {
			coverArt.SetImage(img)
		})
	}()

	if err := app.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
