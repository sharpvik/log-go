package log

import (
	"bufio"
	"io"
	"os"
	"sync"
)

// Log is the central logger struct.
type Log struct {
	lock   sync.Mutex
	level  Priority
	writer *bufio.Writer
	prefix string
}

// New returns a new custom logger instance.
func New(level Priority, writer io.Writer, prefix string) Log {
	return Log{
		level:  level,
		writer: bufio.NewWriter(writer),
		prefix: prefix,
	}
}

// Default returns a default logger instance used withing this package and
// available globally.
func Default() Log {
	return New(LevelError, os.Stdout, "")
}

// WithLevel returns a new logger instance with the specified level.
func (l *Log) WithLevel(level Priority) Log {
	return New(level, l.writer, l.prefix)
}

// WithWriter returns a new logger instance with the specified writer.
func (l *Log) WithWriter(writer io.Writer) Log {
	return New(l.level, writer, l.prefix)
}

// WithPrefix returns a new logger instance with the specified prefix.
func (l *Log) WithPrefix(prefix string) Log {
	return New(l.level, l.writer, prefix)
}

// Panic is equivalent to panic(formattedMessage).
func (l *Log) Panic(format string, args ...interface{}) {
	panic(l.compose(LevelPanic, format, args...))
}

// Fatal is equivalent to Error and os.Exit(1).
func (l *Log) Fatal(format string, args ...interface{}) {
	l.print(LevelFatal, format, args...)
	os.Exit(1)
}

// Error writes formatted error to l.writer.
func (l *Log) Error(format string, args ...interface{}) {
	l.print(LevelError, format, args...)
}

// Warn writes formatted warning to l.writer.
func (l *Log) Warn(format string, args ...interface{}) {
	l.print(LevelWarn, format, args...)
}

// Info writes formatted info message to l.writer.
func (l *Log) Info(format string, args ...interface{}) {
	l.print(LevelInfo, format, args...)
}

// Debug writes formatted debug message to l.writer.
func (l *Log) Debug(format string, args ...interface{}) {
	l.print(LevelDebug, format, args...)
}
