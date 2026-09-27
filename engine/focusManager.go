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

	// снимаем фокус с текущего
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

// GetFocused — возвращает id текущего элемента в фокусе
func (fm *FocusManager) GetFocused() uuid.UUID {
	fm.mu.RLock()
	defer fm.mu.RUnlock()

	if fm.current == -1 || fm.current >= len(fm.order) {
		return uuid.Nil
	}
	return fm.order[fm.current]
}