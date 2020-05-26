package main

import (
	"fmt"
	"log"
	"time"

	"github.com/getlantern/systray"
	"github.com/olehcambel/pomo/sound"
)

type traymenu struct {
	stats      *systray.MenuItem
	toggleMode *systray.MenuItem
}

// func (menu *traymenu) OK() {}

type state struct {
	mode
	cron  *time.Ticker
	menu  traymenu
	timer *timing
	// sound
}

func (s *state) updateLoop() {
	s.updateTitle()
	s.updateModeTitle()
	s.updateStats()

	for {
		<-s.cron.C
		// @see https://golang.org/pkg/time/#hdr-Monotonic_Clocks
		s.timer.start = s.timer.start.Round(0)

		s.updateDeadline()
		s.updateStats()
	}
}

/* STATS */
func (s *state) updateStats() {
	// TODO: check for progress BELOW 0
	s.menu.stats.SetTitle(fmt.Sprintf("%v: %v", tProgress, s.timer.diff()))
}

func (s *state) resetStats() {
	if s.timer.isDeadline {
		s.timer.isDeadline = false
	}
	s.timer.start = time.Now()
	s.updateStats()
}

/* STATS END */

// TODO: maybe MV to mode.go
func (s *state) updateTitle() {
	var title string

	if s.mode == modeWork {
		title = appTitle
	} else {
		title = s.mode.String()
	}

	systray.SetTitle(title)
}

func (s *state) updateModeTitle() {
	// previously ("Finish %v", s.mode)
	s.menu.toggleMode.SetTitle(fmt.Sprint("Switch to ", s.mode.getSwap()))
}

func (s *state) toggleMode() {
	s.mode = s.mode.getSwap()

	s.updateTitle()
	s.updateModeTitle()
	s.resetStats()
}

func (s *state) updateDeadline() {
	if s.mode != modeWork {
		return
	}

	diff := time.Since(s.timer.start).Minutes()
	limit := s.timer.timeLimit.Minutes()

	if s.timer.isDeadline && int(diff)%int(limit) == 0 {
		err := sound.PlayOnce(gSoundPath)

		if err != nil {
			log.Fatal(err)
		}
	} else if diff > limit && !s.timer.isDeadline {
		systray.SetTitle(tDeadline + " " + appTitle)
		s.timer.isDeadline = true

		err := sound.PlayOnce(gSoundPath)
		if err != nil {
			// sentry.CaptureException(err)
			log.Fatal(err)
		}
	}
}

func (s *state) onReady() {
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

func (s *state) onExit() {
	s.cron.Stop()
	log.Print("Exiting")
}
