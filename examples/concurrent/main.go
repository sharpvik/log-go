// This example demonstates that the log-go package is thread-safe.
package main

import (
	"time"

	"github.com/sharpvik/log-go"
)

func init() {
	log.SetLevel(log.LevelDebug)
}

// main spawns two separate threads "one" and "two" and executes another
// function thread "three" after.
func main() {
	go thread("one")
	go thread("two")
	thread("three")
}

// thread logs "Hello from thread {name}" 5 times with an interval of 100ms.
func thread(name string) {
	for i := 0; i < 5; i++ {
		log.Debugf("Hello from thread %s", name)
		time.Sleep(100 * time.Millisecond)
	}
}
