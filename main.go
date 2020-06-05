package main

import (
	"fmt"
	"log"
	"musicPlayer/decoder"
	"musicPlayer/loader"
	"musicPlayer/player"
	"musicPlayer/utils"
)

func main() {

	files := loader.GetMusic()
	randomSong := player.GetSongToPlay(files)
	audio := decoder.GetDecoder(randomSong)
	streamer, format, err := audio.Decode()

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	done := make(chan bool)

	meta, err := utils.GetSongMetaData(randomSong)

	if err != nil {
		fmt.Println("Couldn't get metadata")
	} else {

		fmt.Printf("Title: %s\n", meta.Title)
		fmt.Printf("Artist: %s\n", meta.Artist)
		fmt.Printf("Genre: %s\n", meta.Genre)
	}

	go player.Play(streamer, format, done)
	<-done

	streamer.Close()
}
