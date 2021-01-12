package log

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
var priorityShowValues = [5]string{
	"FATAL",
	"ERROR",
	"WARN ",
	"INFO ",
	"DEBUG",
}

// show function returns Priority flag as a string.
func (p Priority) show() string {
	if int(p) >= len(priorityShowValues) {
		panic("cannot show unknown Priority")
	}
	return priorityShowValues[p]
}
