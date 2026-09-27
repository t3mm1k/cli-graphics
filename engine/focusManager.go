package engine

import (
	"sync"

	"github.com/google/uuid"
)

type FocusManager struct {
	mu      sync.RWMutex
	order   []uuid.UUID // порядок обхода по Tab
	current int         // индекс текущего элемента в order
}

var FocusManagerInstance = &FocusManager{current: -1}

func (fm *FocusManager) Register(id uuid.UUID) {
	fm.mu.Lock()
	defer fm.mu.Unlock()
	fm.order = append(fm.order, id)
}

func (fm *FocusManager) SetFocused(id uuid.UUID) {
	fm.mu.Lock()
	defer fm.mu.Unlock()

	fm.blurCurrent()

	// ищем новый элемент в списке
	for i, compId := range fm.order {
		if compId == id {
			fm.current = i
			fm.focusCurrent()
			return
		}
	}
}


func (fm *FocusManager) GetFocused() uuid.UUID {
	fm.mu.RLock()
	defer fm.mu.RUnlock()

	if fm.current == -1 || fm.current >= len(fm.order) {
		return uuid.Nil
	}
	return fm.order[fm.current]
}

func (fm *FocusManager) FocusNext() {
	fm.mu.Lock()
	defer fm.mu.Unlock()

	if len(fm.order) == 0 {
		return
	}

	fm.blurCurrent()

	if fm.current == -1 {
		fm.current = 0
	} else {
		fm.current = (fm.current + 1) % len(fm.order)
	}

	fm.focusCurrent()
}

func (fm *FocusManager) FocusPrev() {
	fm.mu.Lock()
	defer fm.mu.Unlock()

	if len(fm.order) == 0 {
		return
	}

	fm.blurCurrent()

	if fm.current == -1 {
		fm.current = len(fm.order) - 1
	} else {
		fm.current = (fm.current - 1 + len(fm.order)) % len(fm.order)
	}

	fm.focusCurrent()
}

func (fm *FocusManager) focusCurrent() {
	if fm.current == -1 || fm.current >= len(fm.order) {
		return
	}
	id := fm.order[fm.current]
	comp, exists := Registry.GetComponent(id)
	if !exists {
		return
	}
	if focusable, ok := comp.(Focusable); ok {
		focusable.SetFocus(true)
	}
}

func (fm *FocusManager) blurCurrent() {
	if fm.current == -1 || fm.current >= len(fm.order) {
		return
	}
	id := fm.order[fm.current]
	comp, exists := Registry.GetComponent(id)
	if !exists {
		return
	}
	if focusable, ok := comp.(Focusable); ok {
		focusable.SetFocus(false)
	}
}