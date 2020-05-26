package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	samples := map[string]mode{
		tModeRelax: mode(modeRelax),
		tModeWork:  mode(modeWork),
		tWrongMode: mode(999),
	}

	for name, m := range samples {
		assert.Equal(t, m.String(), name)
	}
}

func TestGetSwap(t *testing.T) {
	samples := map[mode]mode{
		modeRelax: modeWork,
		modeWork:  modeRelax,
		mode(999): mode(999),
	}

	for from, to := range samples {
		assert.Equal(t, to.getSwap(), from)
	}
}
