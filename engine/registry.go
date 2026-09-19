package engine

import (
	"sync"

	"github.com/google/uuid"
)

type ComponentRegistry struct {
	mu         sync.RWMutex
	components map[uuid.UUID]Component
	parents    map[uuid.UUID]uuid.UUID
}

var Registry = &ComponentRegistry{components: make(map[uuid.UUID]Component), parents: make(map[uuid.UUID]uuid.UUID)}

func (registry *ComponentRegistry) AddComponent(component Component) {
	registry.mu.Lock()
	defer registry.mu.Unlock()
	registry.components[component.GetId()] = component
}

func (registry *ComponentRegistry) RemoveComponent(id uuid.UUID) {
	registry.mu.Lock()
	defer registry.mu.Unlock()
	delete(registry.components, id)
}

func (registry *ComponentRegistry) GetComponent(id uuid.UUID) (Component, bool) {
	registry.mu.RLock()
	defer registry.mu.RUnlock()
	component, ok := registry.components[id]
	return component, ok
}

func (registry *ComponentRegistry) SetParent(id uuid.UUID, parent uuid.UUID) {
	registry.mu.Lock()
	defer registry.mu.Unlock()
	registry.parents[id] = parent
}

func (registry *ComponentRegistry) GetParent(id uuid.UUID) (uuid.UUID, bool) {
	registry.mu.RLock()
	defer registry.mu.RUnlock()
	parent, ok := registry.parents[id]
	return parent, ok
}
