package log

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHeader(t *testing.T) {
	location, _ := time.LoadLocation("")
	moment := time.Date(2020, 01, 11, 21, 49, 15, 0, location)

	assert.Equal(t,
		"11/01/2020 21:49:15 \x1b[34mDEBUG\x1b[0m",
		header("", moment, LevelDebug))

	assert.Equal(t,
		"\x1b[36mLOG \x1b[0m11/01/2020 21:49:15 \x1b[34mDEBUG\x1b[0m",
		header("LOG", moment, LevelDebug))
}
