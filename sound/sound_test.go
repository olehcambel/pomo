package sound

import (
	"testing"

	raudio "github.com/olehcambel/pomo/assets/audio"
	"github.com/stretchr/testify/assert"
)

func TestPlayOnce(t *testing.T) {
	err := PlayOnce(raudio.Beep_wav, "wav")

	assert.NoError(t, err)
}
