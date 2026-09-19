package engine

import "github.com/google/uuid"

type Event interface{}

type KeyEvent struct {
	Key string
}

type RerenderEvent struct {
	ComponentId uuid.UUID
}

var EventsQ = make(chan Event, 100)
