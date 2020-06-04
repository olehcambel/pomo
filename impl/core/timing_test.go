package core

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestReset(t *testing.T) {
	now := time.Now()
	timer := Timing{isDeadline: true, start: now}

	assert.Equal(t, now, timer.start)
	timer.reset()

	assert.Equal(t, false, timer.isDeadline)
	assert.NotEqual(t, now, timer.start)
}

func TestFmtDuration(t *testing.T) {
	samples := map[string]time.Duration{
		"1h 0m":    time.Minute * 60,
		"-1m":      time.Minute * -1,
		"2h 59m":   time.Minute * 179,
		"3h 0m":    time.Minute * 180,
		"166h 40m": time.Minute * 10000,
	}

	for expect, d := range samples {
		assert.Equal(t, expect, fmtDuration(d))
	}
}
