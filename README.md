# log-go

Loggin library for Go.

## Log Levels

This library comes with the following log levels:

1. Fatal
2. Error
3. Warn
4. Info
5. Debug

These are enumerated as `LvlDebug`, `LvlInfo`, etc., so that you won't have to
memorise them by numbers.

## Basic Setup

```go
package main

import (
    "os"
    "github.com/sharpvik/log-go" // package log
)

func init() {
    // Change log level.
    log.Level(log.LvlInfo)  // default: LvlError

    // Change log writer.
    file, _ := os.Create("server.log")
    log.Writer(file)        // default: os.Stdout
}

func main() {
    // Computations ...
    x := 40 + 2

    // Print log with priority level Info.
    log.Info("x = %d", x)
    // * 10:45:32 07/10/2020 INFO   x = 42
}
```
