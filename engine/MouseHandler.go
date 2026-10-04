package engine

type MouseHandler interface {
	HandleMouse(mouseEvent *MouseEvent) bool
}
