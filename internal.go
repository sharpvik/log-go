package log

import (
	"fmt"
	"time"

	"github.com/logrusorgru/aurora"
)

// printer is a function that knows its writer, header, and priority level. It
// is simply waiting for the message to print.
type printer func(string, ...interface{})

// header returns a string that represents a log prefix with date, time, and
// log priority level.
func header(prefix string, moment time.Time, level Priority) string {
	var cprefix interface{}
	if len(prefix) == 0 {
		cprefix = ""
	} else {
		cprefix = aurora.Cyan(prefix + " ")
	}
	return fmt.Sprintf("%s%s %s %s",
		cprefix,
		moment.Format("02/01/2006"),
		moment.Format("15:04:05"),
		level.show())
}

// withLevel is a closure that returns a printer function.
func (l *Log) printerWithLevel(level Priority) printer {
	return func(format string, args ...interface{}) {
		l.lock.Lock()
		defer l.lock.Unlock()
		fmt.Fprint(l.writer, header(l.prefix, time.Now(), level), " ")
		fmt.Fprintf(l.writer, format+"\n", args...)
		l.writer.Flush()
	}
}
