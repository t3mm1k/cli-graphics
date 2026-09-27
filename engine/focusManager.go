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