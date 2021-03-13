package log

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHeader(t *testing.T) {
	moment := time.Date(2020, 01, 11, 21, 49, 15, 0, time.Local)
	l := Default()

	assert.Equal(t,
		"11/01/2020 21:49:15 \x1b[34mDEBUG\x1b[0m",
		l.header(moment, LevelDebug))

	l.WithoutColor()
	assert.Equal(t,
		"11/01/2020 21:49:15 DEBUG",
		l.header(moment, LevelDebug))

	l.WithPrefix("LOG")
	assert.Equal(t,
		"LOG 11/01/2020 21:49:15 DEBUG",
		l.header(moment, LevelDebug))

	l.WithColor()
	assert.Equal(t,
		"\x1b[36mLOG \x1b[0m11/01/2020 21:49:15 \x1b[34mDEBUG\x1b[0m",
		l.header(moment, LevelDebug))
}
