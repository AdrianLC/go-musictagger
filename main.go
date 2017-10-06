package main

import (
	id3 "github.com/mikkyang/id3-go"
	"log"
)

func GetGenre(filePath string) string {
	mp3File, err := id3.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	genre := mp3File.Genre()
	defer mp3File.Close()
	return genre
}

func main() {

}
