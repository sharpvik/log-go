package log

import (
	"fmt"
	"time"

	"github.com/logrusorgru/aurora"
)

// header returns a string that contains a log prefix with date, time, and log
// priority level.
func (l *Log) header(moment time.Time, level Priority) string {
	return fmt.Sprintf("%s%s %s %s",
		l.cprefix(),
		moment.Format("02/01/2006"),
		moment.Format("15:04:05"),
		level.show(l.color))
}

// cprefix returns a properly formatted and colored prefix.
func (l *Log) cprefix() interface{} {
	if len(l.prefix) == 0 {
		return ""
	}
	if l.color {
		return aurora.Cyan(l.prefix + " ")
	}
	return l.prefix + " "
}

// compose returns a string ready for printing that contains all the necessary
// components like header and the formatted message.
func (l *Log) compose(
	level Priority, format string, args ...interface{}) string {
	formatted := fmt.Sprintf(format, args...)
	return fmt.Sprintf("%s %s\n",
		l.header(time.Now(), level),
		formatted)
}

// print writes composed record to l.writer.
func (l *Log) print(level Priority, format string, args ...interface{}) {
	if l.level < level {
		return
	}
	l.Lock()
	defer l.Unlock()
	fmt.Fprint(l.writer, l.compose(level, format, args...))
	l.writer.Flush()
}
