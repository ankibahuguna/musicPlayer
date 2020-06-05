package utils

import (
	"fmt"
	"path/filepath"
	"strings"

	id3 "github.com/mikkyang/id3-go"
)

type SongMetaData struct {
	Title  string
	Artist string
	Album  string
	Genre  string
}

func GetSongMetaData(file string) (SongMetaData, error) {
	var songMeta SongMetaData
	mp3, err := id3.Open(file)
	defer mp3.Close()

	if err != nil {
		fmt.Println("Couldn't load metadata")
		return songMeta, err
	}

	songMeta = SongMetaData{
		Title: mp3.Title(), Artist: mp3.Artist(),
		Album: mp3.Album(),
		Genre: mp3.Genre(),
	}

	return songMeta, nil
}

func GetFileName(file string) string {
	mp3, err := id3.Open(file)
	defer mp3.Close()
	f := strings.TrimSuffix(file, filepath.Ext(file))
	if err != nil {
		return f
	}
	name := mp3.Title()

	if name == "" {
		return f
	}

	return name

}
