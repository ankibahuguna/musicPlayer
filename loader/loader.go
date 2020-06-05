package loader

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/h2non/filetype"
	"github.com/h2non/filetype/matchers"
)

func GetMusic() []string {
	var files []string
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	root := fmt.Sprintf("%s/%s", homeDir, "Music")

	e := filepath.Walk(root, visit(&files))
	if e != nil {
		panic(e)
	}

	return files
}

func visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		f, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		head := make([]byte, 261)
		f.Read(head)
		if filetype.IsAudio(head) && matchers.Mp3(head) {
			*files = append(*files, path)
		}
		return nil
	}
}
