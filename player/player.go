package player

import (
	"math/rand"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
)

func Play(streamer beep.Streamer, format beep.Format, done chan bool) {
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))
}

func GetSongToPlay(files []string) string {
	rand.Seed(time.Now().Unix())
	randomSong := files[rand.Intn(len(files))]
	return randomSong
}
