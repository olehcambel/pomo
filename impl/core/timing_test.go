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
