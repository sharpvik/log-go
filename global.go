package log

import "io"

var std = Default()

// SetLevel changes std's priority level.
func SetLevel(level Priority) {
	std.WithLevel(level)
}

// SetWriter changes std's writer.
func SetWriter(writer io.Writer) {
	std.WithWriter(writer)
}

// SetColor changes std's color mode.
func SetColor(color bool) {
	if color {
		std.WithColor()
	} else {
		std.WithoutColor()
	}
}

// Panicf is equivalent to panic(formattedMessage).
func Panicf(format string, args ...interface{}) {
	std.Panicf(format, args...)
}

// Fatalf is equivalent to Error and os.Exit(1).
func Fatalf(format string, args ...interface{}) {
	std.Fatalf(format, args...)
}

// Errorf writes formatted error to l.writer.
func Errorf(format string, args ...interface{}) {
	std.Errorf(format, args...)
}

// Warnf writes formatted warning to l.writer.
func Warnf(format string, args ...interface{}) {
	std.Warnf(format, args...)
}

// Infof writes formatted message to l.writer.
func Infof(format string, args ...interface{}) {
	std.Infof(format, args...)
}

// Debugf writes formatted message to l.writer.
func Debugf(format string, args ...interface{}) {
	std.Debugf(format, args...)
}

// Panic is equivalent to panic(formattedMessage).
func Panic(args ...interface{}) {
	std.Panic(args...)
}

// Fatal is equivalent to Error and os.Exit(1).
func Fatal(args ...interface{}) {
	std.Fatal(args...)
}

// Error writes formatted error to l.writer.
func Error(args ...interface{}) {
	std.Error(args...)
}

// Warn writes formatted warning to l.writer.
func Warn(args ...interface{}) {
	std.Warn(args...)
}

// Info writes formatted info message to l.writer.
func Info(args ...interface{}) {
	std.Info(args...)
}

// Debug writes formatted debug message to l.writer.
func Debug(args ...interface{}) {
	std.Debug(args...)
}
