package logging

type Logger interface {
	Info(msg string, fields ...Field)
	Debug(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Sync() error
}

// Field is a generic key-value field abstraction.
// You can define it as your own type for independence.
type Field struct {
	Key   string
	Value any
}

var Log Logger
