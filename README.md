# log-go

Log-go is a logging library developed with simplicity and thread-safety in mind.
It doesn't support any extra-fancy features. Just like Go itself.

If you wish to improve upon this, you are welcome to send me a Pull Request.

## Features

- Different log levels like **Fatal**, **Error**, **Info**, etc.
- Configurable _buffered_ writer (`os.Stdout`, `*File`, etc.)
- Thread safety
- Prefixes to separate logs from everything else
- Extremely simple setup
- Well-tested

### Coming Soon!

- Colored priority dependent output

## Log Levels

This library comes with the following log levels:

1. Fatal
2. Error
3. Warn
4. Info
5. Debug

These are enumerated as `LvlDebug`, `LvlInfo`, etc., so that you won't have to
memorise them by numbers.

Each `Log` instance has five methods that are named precisely after the levels.
Use them like so:

```go
logger := log.Default()
logger.Warn("this is a warning")
```

## [Basic Setup](examples/basic/main.go)

```go
package main

import (
	"os"
	"github.com/sharpvik/log-go"
)

func init() {
	// Change log level.
	log.SetLevel(log.LvlInfo) // default: LvlError

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
```

## Examples

You are welcome to look at some [examples](examples) too!
