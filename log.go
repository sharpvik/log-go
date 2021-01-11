package log

import (
	"io"
	"os"
)

// Log is the central logger struct.
type Log struct {
	level  Priority
	writer io.Writer
}

// NewLog returns a new custom logger instance.
func NewLog(level Priority, writer io.Writer) Log {
	return Log{level, writer}
}

// DefaultLog returns a default logger instance used withing this package and
// available globally.
func DefaultLog() Log {
	return NewLog(LvlError, os.Stdout)
}

// WithLevel returns a new logger instance with the specified level.
func (l Log) WithLevel(level Priority) Log {
	l.level = level
	return l
}

// WithWriter returns a new logger instance with the specified writer.
func (l Log) WithWriter(writer io.Writer) Log {
	l.writer = writer
	return l
}

// Fatal is equivalent to Error and os.Exit(1).
func (l Log) Fatal(format string, args ...interface{}) {
	l.withLevel(LvlFatal)(format, args...)
	os.Exit(1)
}

// Error writes formatted error to l.writer.
func (l Log) Error(format string, args ...interface{}) {
	l.withLevel(LvlError)(format, args...)
}

// Warn writes formatted warning to l.writer.
func (l Log) Warn(format string, args ...interface{}) {
	l.withLevel(LvlWarn)(format, args...)
}

// Info writes formatted message to l.writer.
func (l Log) Info(format string, args ...interface{}) {
	l.withLevel(LvlInfo)(format, args...)
}

// Debug writes formatted message to l.writer.
func (l Log) Debug(format string, args ...interface{}) {
	l.withLevel(LvlDebug)(format, args...)
}
