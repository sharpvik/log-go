package main

import (
	"errors"

	"github.com/sharpvik/log-go/v2"
)

func init() {
	// Change log level.
	log.SetLevel(log.LevelInfo) // default: LevelError
}

func main() {
	log.Error(errors.New("something went wrong"))
}
