package log

import (
	"fmt"
	"time"

	"github.com/logrusorgru/aurora"
)

// header returns a string that represents a log prefix with date, time, and
// log priority level.
func (l *Log) header(moment time.Time, level Priority) string {
	var cprefix interface{}
	if len(l.prefix) == 0 {
		cprefix = ""
	} else if l.color {
		cprefix = aurora.Cyan(l.prefix + " ")
	} else {
		cprefix = l.prefix + " "
	}
	return fmt.Sprintf("%s%s %s %s",
		cprefix,
		moment.Format("02/01/2006"),
		moment.Format("15:04:05"),
		level.show(l.color))
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
	l.lock.Lock()
	defer l.writer.Flush()
	defer l.lock.Unlock()
	fmt.Fprint(l.writer, l.compose(level, format, args...))
}
