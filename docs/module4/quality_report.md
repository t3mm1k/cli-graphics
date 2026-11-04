> > examples\music_player\main.go:107:24: Error return value of `resp.Body.Close` is not checked (errcheck)

                defer resp.Body.Close()
                                     ^

main.go:32:25: Error return value of `resp.Body.Close` is not checked (errcheck)
defer resp.Body.Close()
^
main.go:54:25: Error return value of `resp.Body.Close` is not checked (errcheck)
defer resp.Body.Close()
^
engine\errors.go:28:5: ST1012: error var NilBufferError should have name of the form ErrFoo (staticcheck)
var NilBufferError = errors.New("engine: target buffer or source data is nil")
^
engine\errors.go:32:5: ST1012: error var NilChildError should have name of the form ErrFoo (staticcheck)
var NilChildError = errors.New("engine: cannot add nil child to container")
^
engine\errors.go:34:5: ST1012: error var CyclicDependencyError should have name of the form ErrFoo (staticcheck)
var CyclicDependencyError = errors.New("engine: cyclic dependency detected: component can`t be child of itself")
^
engine\errors.go:37:5: ST1012: error var NilRootError should have name of the form ErrFoo (staticcheck)
var NilRootError = errors.New("engine: root cannot be nil")
^
engine\errors.go:38:5: ST1012: error var AppAlreadyRunningError should have name of the form ErrFoo (staticcheck)
var AppAlreadyRunningError = errors.New("engine: app is already running")
^
engine\app.go:58:17: func (*App).setTickRate is unused (unused)
func (app *App) setTickRate(tickRate time.Duration) {
^
engine\buffer.go:77:6: func resetTextStyle is unused (unused)
func resetTextStyle(builder \*strings.Builder, style TextStyle) {
^
10 issues:

- errcheck: 3
- staticcheck: 5
- unused: 2
