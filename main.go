package main

import (
	// "log"
	"time"

	"github.com/getlantern/systray"
	"github.com/olehcambel/pomo/impl/core"
	// _ "net/http/pprof"
	// "net/http"
	// _ "net/http"
)

func main() {
	// go func() {
	// 	log.Println(http.ListenAndServe("localhost:6060", nil))
	// }()

	timer := core.NewTiming(time.Minute * 25)
	s := &core.State{Mode: core.ModeRelax, Timer: timer}
	s.Cron = time.NewTicker(time.Minute * 1)

	systray.Run(s.OnReady, s.OnExit)
}
