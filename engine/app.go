package engine

import (
	"cli-graphics/utils"
	"fmt"
	"os"
	"sync"
	"time"

	"golang.org/x/term"
)

type App struct {
	root     Component
	events   chan Event
	tickRate time.Duration

	running  bool
	stopOnce sync.Once
	stopChan chan struct{}

	oldTermState *term.State
}

func NewApp(root Component) *App {
	return &App{
		root:     root,
		events:   make(chan Event, 100),
		tickRate: 500 * time.Millisecond,
		stopChan: make(chan struct{}),
	}
}

func (app *App) setTickRate(tickRate time.Duration) {
	app.tickRate = tickRate
}

func (a *App) Run() error {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return fmt.Errorf("failed to enable raw mode: %w", err)
	}
	a.oldTermState = oldState

	defer a.cleanup()

	fmt.Print("\u001B[?25l")

	go a.listenKeys()

	ticker := time.NewTicker(a.tickRate)
	defer ticker.Stop()

	a.draw()

	for {
		select {
		case <-a.stopChan:
			return nil

		case event := <-a.events:
			switch e := event.(type) {
			case *KeyEvent:
				if e.Key == "Ctrl+C" {
					return nil
				}
				a.handleKey(e.Key)
				a.draw()
			}

		case <-ticker.C:
			if a.root != nil {
				a.root.OnTick()
				a.draw()
			}
		}
	}
}
func (app *App) Stop() {
	app.stopOnce.Do(func() {
		close(app.stopChan)
	})
}
func (a *App) listenKeys() {
	buffer := make([]byte, 16)

	for {
		n, err := os.Stdin.Read(buffer)
		if err != nil {
			return
		}

		key := utils.ParseKey(buffer[:n])

		select {
		case <-a.stopChan:
			return
		case a.events <- &KeyEvent{Key: key}:

		}
	}
}

func (a *App) handleKey(key string) {
	if a.root == nil {
		return
	}
	if focusable, ok := a.root.(Focusable); ok {
		if key == "Tab" {
			if !focusable.HandleKey("Tab") {
				focusable.SetFocus(true)
			}
		} else {
			focusable.HandleKey(key)
		}
	}
}

func (a *App) cleanup() {
	if a.oldTermState != nil {
		_ = term.Restore(int(os.Stdin.Fd()), a.oldTermState)
	}

	fmt.Print("\u001B[?25h\u001B[2J\u001B[H")
}

func (a *App) draw() {
	if a.root == nil {
		return
	}
	a.root.Render()
	if base, ok := a.root.(interface{ Flush() }); ok {
		base.Flush()
	} else if bc, ok := a.root.(Container); ok {
		bc.Flush()
	}
}
