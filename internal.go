package log

import (
	"fmt"
	"time"
)

// printer is a function that knows its writer, header, and priority level. It
// is simply waiting for the message to print.
type printer func(string, ...interface{})

// header returns a string that represents a log prefix with date, time, and
// log priority level.
func header(moment time.Time, level Priority) string {
	return fmt.Sprintf("%s %s %s",
		moment.Format("02/01/2006"),
		moment.Format("15:04:05"),
		level.Show())
}

// withLevel is a closure that returns a printer function.
func (l Log) withLevel(level Priority) printer {
	return func(format string, args ...interface{}) {
		fmt.Fprint(l.writer, header(time.Now(), level), " ")
		fmt.Fprintf(l.writer, format+"\n", args...)
	}
}
