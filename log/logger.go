package log

type Logger struct {
	Log ILog `inject:""`
}

type ILog interface {
	Log(e *Entry) error
}
