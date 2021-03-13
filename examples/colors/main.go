// This example shows you all colours of the rainbow! :)
package main

import "github.com/sharpvik/log-go"

func init() {
	log.SetLevel(log.LevelDebug)
	log.SetColor(false)
	log.Info("you can now log without color if you want!")
	log.SetColor(true)
}

func main() {
	log.Debugf("launching nuclear missiles...")
	log.Infof("all warheads left their stations")
	log.Warnf("the targets will be reached soon")
	log.Errorf("targets have not been destroyed")
	log.Fatalf("looks like we've been hacked :(")
}
