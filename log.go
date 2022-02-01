package log

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sync"
)

// Log is the central logger struct.
type Log struct {
	sync.Mutex
	level  Priority
	writer *bufio.Writer
	prefix string
	color  bool
}

// New returns a new custom logger instance.
func New(level Priority, writer io.Writer, prefix string, color bool) *Log {
	return &Log{
		level:  level,
		writer: bufio.NewWriter(writer),
		prefix: prefix,
		color:  true,
	}
}

// Default returns a default logger instance used withing this package and
// available globally.
func Default() *Log {
	return New(LevelError, os.Stdout, "", true)
}

// WithLevel returns a new logger instance with the specified level.
func (l *Log) WithLevel(level Priority) *Log {
	l.level = level
	return l
}

// WithWriter returns a new logger instance with the specified writer.
func (l *Log) WithWriter(writer io.Writer) *Log {
	l.writer = bufio.NewWriter(writer)
	return l
}

// WithPrefix returns a new logger instance with the specified prefix.
func (l *Log) WithPrefix(prefix string) *Log {
	l.prefix = prefix
	return l
}

// WithColor returns a new logger instance with color mode turned on.
func (l *Log) WithColor() *Log {
	l.color = true
	return l
}

// WithoutColor returns a new logger instance with color mode turned off.
func (l *Log) WithoutColor() *Log {
	l.color = false
	return l
}

// Panicf is equivalent to panic(formattedMessage).
func (l *Log) Panicf(format string, args ...interface{}) {
	panic(l.compose(LevelPanic, format, args...))
}

// Fatalf is equivalent to Error and os.Exit(1).
func (l *Log) Fatalf(format string, args ...interface{}) {
	l.print(LevelFatal, format, args...)
	os.Exit(1)
}

// Errorf writes formatted error to l.writer.
func (l *Log) Errorf(format string, args ...interface{}) {
	l.print(LevelError, format, args...)
}

// Warnf writes formatted warning to l.writer.
func (l *Log) Warnf(format string, args ...interface{}) {
	l.print(LevelWarn, format, args...)
}

// Infof writes formatted info message to l.writer.
func (l *Log) Infof(format string, args ...interface{}) {
	l.print(LevelInfo, format, args...)
}

// Debugf writes formatted debug message to l.writer.
func (l *Log) Debugf(format string, args ...interface{}) {
	l.print(LevelDebug, format, args...)
}

// Panic is equivalent to panic(formattedMessage).
func (l *Log) Panic(args ...interface{}) {
	panic(l.compose(LevelPanic, fmt.Sprint(args...)))
}

// Fatal is equivalent to Error and os.Exit(1).
func (l *Log) Fatal(args ...interface{}) {
	l.print(LevelFatal, fmt.Sprint(args...))
	os.Exit(1)
}

// Error writes formatted error to l.writer.
func (l *Log) Error(args ...interface{}) {
	l.print(LevelError, fmt.Sprint(args...))
}

// Warn writes formatted warning to l.writer.
func (l *Log) Warn(args ...interface{}) {
	l.print(LevelWarn, fmt.Sprint(args...))
}

// Info writes formatted info message to l.writer.
func (l *Log) Info(args ...interface{}) {
	l.print(LevelInfo, fmt.Sprint(args...))
}

// Debug writes formatted debug message to l.writer.
func (l *Log) Debug(args ...interface{}) {
	l.print(LevelDebug, fmt.Sprint(args...))
}
