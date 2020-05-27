// Package sound provides funcs to work with audio formats
package sound

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

// ReadSeekCloser is io.Closer and io.ReadSeeker.
type ReadSeekCloser interface {
	io.ReadSeeker
	io.Closer
}

type bytesNewReadCloser struct {
	reader *bytes.Reader
}

func (b *bytesNewReadCloser) Read(buf []byte) (int, error) {
	return b.reader.Read(buf)
}

func (b *bytesNewReadCloser) Seek(offset int64, whence int) (int64, error) {
	return b.reader.Seek(offset, whence)
}

func (b *bytesNewReadCloser) Close() error {
	b.reader = nil
	return nil
}

// BytesNewReadCloser creates ReadSeekCloser from bytes.
func BytesNewReadCloser(b []byte) ReadSeekCloser {
	return &bytesNewReadCloser{reader: bytes.NewReader(b)}
}

// PlayOnce opens audio file, plays it once and then close
func PlayOnce(b []byte, fileExt string) error {
	var stream beep.StreamSeekCloser
	var format beep.Format
	var err error

	switch fileExt {
	case "mp3":
		stream, format, err = mp3.Decode(BytesNewReadCloser(b))
	case "wav":
		stream, format, err = wav.Decode(bytes.NewReader(b))
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
