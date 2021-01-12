// This example shows you all colours of the rainbow! :)
package main

import "github.com/sharpvik/log-go"

func init() {
	log.SetLevel(log.LevelDebug)
}

func main() {
	log.Debug("launching nuclear missiles...")
	log.Info("all warheads left their stations")
	log.Warn("the targets will be reches soon")
	log.Error("targets have not been destroyed")
	log.Fatal("looks like we've been hacked :(")
}
