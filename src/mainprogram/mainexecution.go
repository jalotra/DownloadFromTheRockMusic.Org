package main

import (
	"downloadfiles"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: download url filename")
		os.Exit(1)
	}
	url := os.Args[1]
	filename := os.Args[2]

	// Changing filename to filepath so that the songs are downloaded
	// in separate folder Songs.

	os.Chdir("../")
	os.Mkdir("songs", 777)
	os.Chdir("./songs")

	err := downloadfiles.DownloadFile(url, filename)
	if err != nil {
		panic(err)
	}

}
