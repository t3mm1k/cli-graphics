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

	enableLogging bool
	logFile       *os.File
}

func NewApp(root Component) *App {
	return &App{
		root:     root,
		events:   make(chan Event, 100),
		tickRate: 500 * time.Millisecond,
		stopChan: make(chan struct{}),
	}
}

// SetLogging включает или выключает запись логов в файл app.log
func (app *App) SetLogging(enable bool) *App {
	app.enableLogging = enable
	return app
}

func (app *App) setTickRate(tickRate time.Duration) {
	app.tickRate = tickRate
}

func (a *App) Run() error {
	if a.enableLogging {
		f, err := enableFileLogging("app.log")
		if err != nil {
			return fmt.Errorf("failed to open app.log: %w", err)
		}
		a.logFile = f
		Log.Info("App started with file logging", "tickRate", a.tickRate)
	}

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		Log.Error("Failed to enable raw mode", "err", err)
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

	if a.logFile != nil {
		Log.Info("App cleanup completed, closing log file")
		_ = a.logFile.Close()
		resetLogger()
		a.logFile = nil
	}
}

func (a *App) draw() {
	if a.root == nil {
		return
	}
	a.root.Render()
	if flusher, ok := a.root.(interface{ Flush() }); ok {
		flusher.Flush()
	}
}
