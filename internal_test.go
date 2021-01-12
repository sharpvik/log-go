package log

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHeader(t *testing.T) {
	location, _ := time.LoadLocation("")
	moment := time.Date(2020, 01, 11, 21, 49, 15, 0, location)
	assert.Equal(t, "11/01/2020 21:49:15 DEBUG", header(moment, LevelDebug))
}
