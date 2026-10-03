//go:build !windows

package engine

func setupTerminal() func() {
	return func() {}
}
