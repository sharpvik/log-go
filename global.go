package log

import "io"

var std = DefaultLog()

// WithLevel returns a new logger instance with the specified level.
func WithLevel(level Priority) {
	std = std.WithLevel(level)
}

// WithWriter returns a new logger instance with the specified writer.
func WithWriter(writer io.Writer) {
	std = std.WithWriter(writer)
}

// Fatal is equivalent to Error and os.Exit(1).
var Fatal = std.Fatal

// Error writes formatted error to l.writer.
var Error = std.Fatal

// Warn writes formatted warning to l.writer.
var Warn = std.Fatal

// Info writes formatted message to l.writer.
var Info = std.Fatal

// Debug writes formatted message to l.writer.
var Debug = std.Fatal
