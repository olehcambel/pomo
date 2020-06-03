package core

import (
	"fmt"
	"log"
	"time"

	"github.com/getlantern/systray"
	raudio "github.com/olehcambel/pomo/assets/audio"
	"github.com/olehcambel/pomo/sound"
)

type traymenu struct {
	stats      *systray.MenuItem
	toggleMode *systray.MenuItem
}

// func (menu *traymenu) OK() {}

// State of the app
type State struct {
	Mode  mode
	Cron  *time.Ticker
	menu  traymenu
	Timer *Timing
	// sound
}

func (s *State) updateLoop() {
	s.updateTitle()
	s.updateModeTitle()
	s.updateStats()
	// FIXME: ALSA lib pcm_dmix.c:1089:(snd_pcm_dmix_open) unable to open slave

	for range s.Cron.C {
		// @see https://golang.org/pkg/time/#hdr-Monotonic_Clocks
		s.Timer.start = s.Timer.start.Round(0)

		s.updateDeadline()
		s.updateStats()
	}
	// for {
	// 	<-s.Cron.C
	// 	// @see https://golang.org/pkg/time/#hdr-Monotonic_Clocks
	// 	s.Timer.start = s.Timer.start.Round(0)

	// 	s.updateDeadline()
	// 	s.updateStats()
	// }
}

/* STATS */
func (s *State) updateStats() {
	// TODO: check for progress BELOW 0
	s.menu.stats.SetTitle(fmt.Sprintf("%v: %v", tProgress, s.Timer.diff()))
}

func (s *State) resetStats() {
	if s.Timer.isDeadline {
		s.Timer.isDeadline = false
	}
	s.Timer.start = time.Now()
	s.updateStats()
}

/* STATS END */

// TODO: maybe MV to mode.go
func (s *State) updateTitle() {
	var title string

	if s.Mode == ModeWork {
		title = appTitle
	} else {
		title = s.Mode.String()
	}

	// Debug(title)
	// FIXME: after long run, systray.SetTitle is not updating title
	systray.SetTitle(title)
}

func (s *State) updateModeTitle() {
	// previously ("Finish %v", s.Mode)
	s.menu.toggleMode.SetTitle(fmt.Sprint("Switch to ", s.Mode.getSwap()))
}

func (s *State) toggleMode() {
	s.Mode = s.Mode.getSwap()

	s.updateTitle()
	s.updateModeTitle()
	s.resetStats()
}

func (s *State) updateDeadline() {
	if s.Mode != ModeWork {
		return
	}

	diff := time.Since(s.Timer.start).Minutes()
	limit := s.Timer.timeLimit.Minutes()

	if s.Timer.isDeadline && int(diff)%int(limit) == 0 {
		err := sound.PlayOnce(raudio.Beep_wav, gWavExt)

		if err != nil {
			log.Fatal(err)
		}
	} else if diff > limit && !s.Timer.isDeadline {
		systray.SetTitle(tDeadline + " " + appTitle)
		s.Timer.isDeadline = true

		err := sound.PlayOnce(raudio.Beep_wav, gWavExt)
		if err != nil {
			// sentry.CaptureException(err)
			log.Fatal(err)
		}
	}
}

// OnReady runs when systray is ready
func (s *State) OnReady() {
	s.menu.stats = systray.AddMenuItem("---", "Stats")
	s.menu.stats.Disable()

	systray.AddSeparator()

	s.menu.toggleMode = systray.AddMenuItem("---", "Toggle mode")

	systray.AddSeparator()

	mQuit := systray.AddMenuItem(tQuit, "Quit the app")

	go s.updateLoop()
	// eventHandler
	go func() {
		for {
			select {
			case <-s.menu.toggleMode.ClickedCh:
				s.toggleMode()
			case <-mQuit.ClickedCh:
				systray.Quit()
			}
		}
	}()
}

// OnExit runs when systray is going to exit
func (s *State) OnExit() {
	s.Cron.Stop()
	log.Print("Exiting")
}
