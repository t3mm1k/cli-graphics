package engine

type Event interface{}

type KeyEvent struct {
	Key string
}

var EventsQ = make(chan Event, 100)
