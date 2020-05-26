package sound

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayOnce(t *testing.T) {
	err := PlayOnce("../assets/sounds/beep.wav")

	assert.NoError(t, err)
}
