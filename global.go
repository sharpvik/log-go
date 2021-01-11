package log

import "io"

var std = DefaultLog()

// SetLevel returns a new logger instance with the specified level.
func SetLevel(level Priority) {
	std = std.WithLevel(level)
}

// SetWriter returns a new logger instance with the specified writer.
func SetWriter(writer io.Writer) {
	std = std.WithWriter(writer)
}

// Fatal is equivalent to Error and os.Exit(1).
var Fatal = std.Fatal

// Error writes formatted error to l.writer.
var Error = std.Error

// Warn writes formatted warning to l.writer.
var Warn = std.Warn

// Info writes formatted message to l.writer.
var Info = std.Info

// Debug writes formatted message to l.writer.
var Debug = std.Debug
