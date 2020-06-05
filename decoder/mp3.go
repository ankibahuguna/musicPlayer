package decoder

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/wav"
)

type Mp3 struct {
	File io.ReadCloser
}

type Wav struct {
	File io.Reader
}

type Flac struct {
	File io.ReadCloser
}

type Audio interface {
	Decode() (beep.StreamSeekCloser, beep.Format, error)
}

func (m Mp3) Decode() (beep.StreamSeekCloser, beep.Format, error) {
	streamer, format, err := mp3.Decode(m.File)

	return streamer, format, err
}

func (f Flac) Decode() (beep.StreamSeekCloser, beep.Format, error) {
	streamer, format, err := mp3.Decode(f.File)

	return streamer, format, err
}

func (w Wav) Decode() (beep.StreamSeekCloser, beep.Format, error) {
	streamer, format, err := wav.Decode(w.File)

	return streamer, format, err
}

func GetDecoder(song string) Audio {
	file, err := os.Open(song)

	if err != nil {
		log.Fatal(err)
	}

	fileType := filepath.Ext(song)

	var audio Audio

	switch fileType {
	case ".mp3":
		audio = Mp3{File: file}
	case ".wav":
		audio = Wav{File: file}
	case ".flac":
		audio = Flac{File: file}
	default:
		err := fmt.Sprintf("%s:%s", "Invalid file type ", fileType)
		log.Fatal(err)
	}

	return audio
}
