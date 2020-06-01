package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	samples := map[string]mode{
		tModeRelax: mode(ModeRelax),
		tModeWork:  mode(ModeWork),
		tWrongMode: mode(999),
	}

	for name, m := range samples {
		assert.Equal(t, m.String(), name)
	}
}

func TestGetSwap(t *testing.T) {
	samples := map[mode]mode{
		ModeRelax: ModeWork,
		ModeWork:  ModeRelax,
		mode(999): mode(999),
	}

	for from, to := range samples {
		assert.Equal(t, to.getSwap(), from)
	}
}
