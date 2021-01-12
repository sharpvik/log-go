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

// priorityShowValues is a map that connects Priority with its string
// representation. Notice, that all strings here are exactly 5 chars long so as
// to be perfectly aligned when printed.
var priorityShowValues = map[Priority]string{
	LevelFatal: "FATAL",
	LevelError: "ERROR",
	LevelWarn:  "WARN ",
	LevelInfo:  "INFO ",
	LevelDebug: "DEBUG",
}

// show function returns Priority flag as a string.
func (p Priority) show() (str string) {
	str, ok := priorityShowValues[p]
	if !ok {
		panic("cannot show unknown Priority")
	}
	return
}
