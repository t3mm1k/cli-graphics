package engine

import (
	"bytes"
	"cli-graphics/utils"
	"fmt"
	"os"
	"runtime/debug"
	"sync"
	"time"

	"golang.org/x/term"
)

type App struct {
	buffer   *Buffer
	root     Component
	events   chan Event
	actions  chan func()
	tickRate time.Duration

	shortcuts map[string]func()

	running  bool
	stopOnce sync.Once
	stopChan chan struct{}

	oldTermState *term.State

	enableLogging bool
	logFile       *os.File
}

func NewApp(root Component) *App {
	buf, err := NewBuffer(80, 24)
	if err != nil {
		panic(err)
	}
	return &App{
		buffer:    buf,
		root:      root,
		events:    make(chan Event, 100),
		tickRate:  500 * time.Millisecond,
		stopChan:  make(chan struct{}),
		actions:   make(chan func(), 100),
		shortcuts: make(map[string]func()),
	}
}

func (app *App) SetLogging(enable bool) *App {
	app.enableLogging = enable
	return app
}

func (app *App) setTickRate(tickRate time.Duration) {
	app.tickRate = tickRate
}

func (a *App) Run() error {
	if a.root == nil {
		return NilRootError
	}
	if a.running {
		return AppAlreadyRunningError
	}

	if a.enableLogging {
		f, err := enableFileLogging("app.log")
		if err != nil {
			return fmt.Errorf("app: failed to open app.log: %w", err)
		}
		a.logFile = f
	}

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		Log.Error("app: failed to enable raw mode", "err", err)
		return fmt.Errorf("app: failed to enable raw mode: %w", err)
	}
	a.oldTermState = oldState

	defer func() {
		a.cleanup()
		if r := recover(); r != nil {
			Log.Error("Application panicked", "error", r, "stack", string(debug.Stack()))
			panic(r)
		}
	}()
	defer setupTerminal()()

	fmt.Print("\u001B[?1049h\u001B[?25l\u001B[?7l\033[?1000h\033[?1002h\033[?1006h")

	go a.listenKeys()
	go a.listenResize()

	ticker := time.NewTicker(a.tickRate)
	defer ticker.Stop()

	a.running = true

	if w, h, err := term.GetSize(int(os.Stdout.Fd())); err == nil && w > 0 && h > 0 {
		a.buffer.SetSize(w, h)
		a.root.SetSize(w, h)
	}
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
			case *MouseEvent:
				Log.Debug("app: mouse event", "event", e)

			case *TerminalResizeEvent:
				a.buffer.SetSize(e.Width, e.Height)
				a.root.SetSize(e.Width, e.Height)
				a.draw()
			}

		case action := <-a.actions:
			action()
			a.draw()

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

func (a *App) listenResize() {
	ticker := time.NewTicker(150 * time.Millisecond)
	defer ticker.Stop()
	lastW, lastH := a.buffer.GetSize()
	for {
		select {
		case <-a.stopChan:
			return
		case <-ticker.C:
			w, h, err := term.GetSize(int(os.Stdout.Fd()))
			//h -= 1
			if err != nil || w <= 0 || h <= 0 {
				continue
			}
			if w != lastW || h != lastH {
				lastW, lastH = w, h
				select {
				case <-a.stopChan:
					return
				case a.events <- &TerminalResizeEvent{Width: w, Height: h}:
				}
			}
		}
	}
}

func (a *App) listenKeys() {
	buffer := make([]byte, 64)

	for {
		n, err := os.Stdin.Read(buffer)
		if err != nil {
			return
		}

		if bytes.HasPrefix(buffer[:n], []byte("\x1b[<")) {
			if mouse, ok := utils.ParseSGRMouse(buffer[:n]); ok {
				select {
				case <-a.stopChan:
					return
				case a.events <- &MouseEvent{X: mouse.X, Y: mouse.Y, MouseButton: mouse.Button, MouseAction: mouse.Action}:
				}
			}
		} else {
			key := utils.ParseKey(buffer[:n])
			select {
			case <-a.stopChan:
				return
			case a.events <- &KeyEvent{Key: key}:

			}
		}
	}
}

func (a *App) handleKey(key string) {
	if a.root == nil {
		return
	}

	if key == "Tab" {
		FocusManagerInstance.FocusNext()
		return
	}
	if key == "Shift+Tab" {
		FocusManagerInstance.FocusPrev()
		return
	}

	if focusable, ok := a.root.(Focusable); ok {
		if !focusable.HandleKey(key) {
			if handler, exists := a.shortcuts[key]; exists && handler != nil {
				handler()
			}
		}
	}
}

func (a *App) AddShortcut(shortcut string, handler func()) {
	a.shortcuts[shortcut] = handler
}

func (a *App) cleanup() {
	if a.oldTermState != nil {
		_ = term.Restore(int(os.Stdin.Fd()), a.oldTermState)
	}

	fmt.Print("\u001B[?7h\u001B[?25h\u001B[?1049l\\033[?1006l\\033[?1002l\\033[?1000l")

	if a.logFile != nil {
		_ = a.logFile.Close()
		resetLogger()
		a.logFile = nil
	}
}

func (a *App) draw() {
	if a.root == nil || a.buffer == nil {
		return
	}
	a.buffer.Clear()
	canvas := NewCanvas(a.buffer)
	a.root.Render(canvas)
	// a.buffer.Flush()
	fmt.Print("\033[H" + a.buffer.Flush())
}

func (a *App) Post(action func()) {
	select {
	case <-a.stopChan:
		return
	case a.actions <- action:
	}
}
