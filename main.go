package main

import (
	"cli-graphics/engine"
	"cli-graphics/widgets"
	"fmt"
	"image"
	"net/http"

	_ "image/jpeg"
	_ "image/png"
)

func main() {
	rootScreen := widgets.NewBox(0, 0, 80, 24)
	btnImage := widgets.NewButton(10, 20, "Image", nil)
	picture := widgets.NewImage(40, 10, 40, 20, nil)
	rootScreen.AddChild(btnImage)
	rootScreen.AddChild(picture)
	rootScreen.SetFocus(true)

	app := engine.NewApp(rootScreen)
	app.SetLogging(true)

	btnImage.OnClick = func() {
		go func() {
			imgURL := "https://avatars.yandex.net/get-music-content/16334817/73d15c0c.a.16218390-2/400x400"
			resp, err := http.Get(imgURL)
			if err != nil {
				return
			}
			defer resp.Body.Close()
			if resp.StatusCode != http.StatusOK {
				return
			}
			img, _, err := image.Decode(resp.Body)
			if err != nil {
				return
			}

			app.Post(func() {
				picture.SetImage(img)
			})
		}()
	}

	if err := app.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
