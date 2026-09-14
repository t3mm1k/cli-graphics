package engine

import (
	"sync"

	"github.com/google/uuid"
)

type ComponentRegistry struct {
	mu         sync.RWMutex
	components map[uuid.UUID]Component
}

var Registry = &ComponentRegistry{components: make(map[uuid.UUID]Component)}

func (registry *ComponentRegistry) Add(component Component) {
	registry.mu.Lock()
	defer registry.mu.Unlock()
	registry.components[component.GetId()] = component
}

func (registry *ComponentRegistry) Remove(id uuid.UUID) {
	registry.mu.Lock()
	defer registry.mu.Unlock()
	delete(registry.components, id)
}

func (registry *ComponentRegistry) GetWidget(id uuid.UUID) (Component, bool) {
	registry.mu.RLock()
	defer registry.mu.RUnlock()
	component, ok := registry.components[id]
	return component, ok
}
