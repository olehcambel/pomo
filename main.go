package main

import (
	"time"

	"github.com/getlantern/systray"
)

func main() {
	timer := newTiming(time.Minute * 25)
	s := state{mode: modeRelax, timer: timer}
	s.cron = time.NewTicker(time.Minute * 1)

	systray.Run(s.onReady, s.onExit)
}
