package log

import "github.com/logrusorgru/aurora"

// Priority is a type alias for logging priority levels.
type Priority uint8

// Log Priority levels are listed here as constants for convenience.
const (
	LevelFatal Priority = iota
	LevelError
	LevelWarn
	LevelInfo
	LevelDebug
)

// priorityShowValues is an array that Priority string representations.
// Notice, that all strings here are exactly 5 chars long so as to be perfectly
// aligned when printed.
var priorityShowValues = [5]struct {
	str string
	clr func(interface{}) aurora.Value
}{
	{"FATAL", aurora.Magenta},
	{"ERROR", aurora.Red},
	{"WARN ", aurora.Yellow},
	{"INFO ", aurora.Green},
	{"DEBUG", aurora.Blue},
}

// show function returns Priority flag as a string.
func (p Priority) show() aurora.Value {
	if int(p) >= len(priorityShowValues) {
		panic("cannot show unknown Priority")
	}
	value := priorityShowValues[p]
	return value.clr(value.str)
}
