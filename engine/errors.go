package engine

import (
	"errors"
	"fmt"
)

//Ошибки буффера

type InvalidBufferSizeError struct {
	W, H int
}

func (e InvalidBufferSizeError) Error() string {
	return fmt.Sprintf("engine: invalid buffer size (%d) of size (%d)", e.W, e.H)
}

type OutOfBoundsError struct {
	X, Y, W, H       int
	BufferW, BufferH int
}

func (e *OutOfBoundsError) Error() string {
	return fmt.Sprintf("engine: cannot draw region (%d,%d %dx%d): exceeds buffer size (%dx%d)",
		e.X, e.Y, e.W, e.H, e.BufferW, e.BufferH)
}

var NilBufferError = errors.New("engine: target buffer or source data is nil")

//Ошибки компонентов

var NilChildError = errors.New("engine: cannot add nil child to container")

var CyclicDependencyError = errors.New("engine: cyclic dependency detected: component can`t be child of itself")

// Ошибки Приложения
var NilRootError = errors.New("engine: root cannot be nil")
var AppAlreadyRunningError = errors.New("engine: app is already running")
