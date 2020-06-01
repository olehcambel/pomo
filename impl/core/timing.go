package core

// package Timing

import (
	"fmt"
	"time"
)

// Timing is a struct for time duration
type Timing struct {
	timeLimit  time.Duration
	start      time.Time
	isDeadline bool // initial
}

// NewTiming creates new instance of Timing. MV to Timing lib
func NewTiming(timeLimit time.Duration) *Timing {
	return &Timing{
		timeLimit: timeLimit,
		start:     time.Now(),
	}
}

func (t *Timing) diff() string {
	return fmtDuration(time.Since(t.start))
}

func (t *Timing) reset() {
	if t.isDeadline {
		t.isDeadline = false
	}

	t.start = time.Now()
}

func fmtDuration(d time.Duration) string {
	// d could be negative (-n)

	if m := d.Minutes(); m < 60 {
		return fmt.Sprintf("%1.fm", m)
	}

	h := d.Minutes() / 60
	m := d.Minutes() - float64(int(h)*60)
	return fmt.Sprintf("%1.fh %1.fm", h, m)
}
