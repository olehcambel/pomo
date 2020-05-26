package main

// package timing

import (
	"fmt"
	"time"
)

type timing struct {
	timeLimit  time.Duration
	start      time.Time
	isDeadline bool // initial
}

func newTiming(timeLimit time.Duration) *timing {
	return &timing{
		timeLimit: timeLimit,
		start:     time.Now(),
	}
}

func (t *timing) diff() string {
	return fmtDuration(time.Since(t.start))
}

func (t *timing) reset() {
	if t.isDeadline {
		t.isDeadline = false
	}

	t.start = time.Now()
}

func fmtDuration(d time.Duration) string {
	// d could be "-n"
	m := d / time.Minute

	// d -= m * time.Minute
	// s := d / time.Second
	// return fmt.Sprintf("%dm %ds", int(m), int(s))

	return fmt.Sprintf("%dm", int(m))
}
