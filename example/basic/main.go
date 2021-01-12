package main

import (
	"os"

	"github.com/sharpvik/log-go"
)

func init() {
	// Change log level.
	log.SetLevel(log.LevelInfo) // default: LevelError

	// Change log writer.
	file, _ := os.Create("server.log")
	log.SetWriter(file) // default: os.Stdout
}

func main() {
	// Computations ...
	x := 40 + 2

	// Print log with priority level Info.
	log.Info("x = %d", x)
	// * 11/01/2021 23:07:08 INFO  x = 42
}
