package engine

type Event interface{}

type KeyEvent struct {
	Key string
}

type RerenderEvent struct {
}
