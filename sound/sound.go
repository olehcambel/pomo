// Package sound provides funcs to work with audio formats
package sound

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

// PlayOnce opens audio file, plays it once and then close
func PlayOnce(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	var stream beep.StreamSeekCloser
	var format beep.Format

	switch fileExt := filepath.Ext((path)); fileExt {
	case ".mp3":
		stream, format, err = mp3.Decode(f)
	case ".wav":
		stream, format, err = wav.Decode(f)
	default:
		return fmt.Errorf("Format not available: %v", fileExt)
	}
	if err != nil {
		return err
	}

	defer stream.Close()
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(stream, beep.Callback(func() {
		done <- true
	})))
	<-done
	close(done)

	return nil
}
