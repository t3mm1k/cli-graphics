package engine

type Event interface{}

type KeyEvent struct {
	key string
}

type RerenderEvent struct {
}
