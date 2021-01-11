package log

// Priority is a type alias for logging priority levels.
type Priority uint8

// Show function returns Priority flag as a string.
func (p Priority) Show() string {
	switch p {
	case LvlFatal:
		return "FATAL"
	case LvlError:
		return "ERROR"
	case LvlWarn:
		return "WARN "
	case LvlInfo:
		return "INFO "
	case LvlDebug:
		return "DEBUG"
	}
	panic("control must never reach this line; make sure that the switch is exhaustive")
}
