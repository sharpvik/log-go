package log

import "io"

var std = Default()

// SetLevel changes std's priority level.
func SetLevel(level Priority) {
	std = std.WithLevel(level)
}

// SetWriter changes std's writer.
func SetWriter(writer io.Writer) {
	std = std.WithWriter(writer)
}

// Fatal is equivalent to Error and os.Exit(1).
func Fatal(format string, args ...interface{}) {
	std.Fatal(format, args...)
}

// Error writes formatted error to l.writer.
func Error(format string, args ...interface{}) {
	std.Error(format, args...)
}

// Warn writes formatted warning to l.writer.
func Warn(format string, args ...interface{}) {
	std.Warn(format, args...)
}

// Info writes formatted message to l.writer.
func Info(format string, args ...interface{}) {
	std.Info(format, args...)
}

// Debug writes formatted message to l.writer.
func Debug(format string, args ...interface{}) {
	std.Debug(format, args...)
}
