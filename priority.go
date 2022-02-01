package log

import "github.com/logrusorgru/aurora"

// Priority is a type alias for logging priority levels.
type Priority uint8

// Log Priority levels are listed here as constants for convenience.
const (
	LevelPanic Priority = iota
	LevelFatal
	LevelError
	LevelWarn
	LevelInfo
	LevelDebug
)

// priorityShowValues is an array that contains Priority string representations.
// Notice, that all strings here are exactly 5 chars long so as to be perfectly
// aligned when printed.
var priorityShowValues = [6]struct {
	text  string
	color func(interface{}) aurora.Value
}{
	{"PANIC", aurora.Magenta},
	{"FATAL", aurora.Magenta},
	{"ERROR", aurora.Red},
	{"WARN ", aurora.Yellow},
	{"INFO ", aurora.Green},
	{"DEBUG", aurora.Blue},
}

// show function returns Priority flag as a string colored depending on the
// color argument passed in.
func (p Priority) show(color bool) interface{} {
	if int(p) >= len(priorityShowValues) {
		panic("cannot show unknown Priority")
	}
	value := priorityShowValues[p]
	if color {
		return value.color(value.text)
	}
	return value.text
}
