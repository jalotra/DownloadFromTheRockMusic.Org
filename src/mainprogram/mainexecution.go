package main

import (
	"downloadfiles"
	"fmt"
	"geturldata"
	"os"
	"parsethemusicmp3snames"
)

func parseWEBSITEADDRESS() string {

	if len(os.Args) != 3 {
		fmt.Println("usage: ./mainexecution therockmusic.orgAlbumURL foldername")
		os.Exit(1)
	}
	// var Reader io.Reader
	documentTree := geturldata.ReturnURLBODY(geturldata.DownloadWEBADDRESS())

	return documentTree

}

func gettingAlbumTitles() {
	parsethemusicmp3snames.ReturnAnchorTags(parsethemusicmp3snames.HTMLParser(parseWEBSITEADDRESS()))

}

func downloaderASSEMBLY() {
	if len(os.Args) != 3 {
		fmt.Println("usage: ./mainexecution therockmusic.orgAlbumURL foldername")
		os.Exit(1)
	}
	url := os.Args[1]
	filename := os.Args[2]

	// Changing filename to filepath so that the songs are downloaded
	// in separate folder Songs and then in the albums folder .

	os.Chdir("../")
	os.Mkdir("songs", 777)
	os.Chdir("./songs")
	os.Mkdir(geturldata.ParseWEBSITEADDRESS(os.Args[1]), 777)
	os.Chdir("./" + geturldata.ParseWEBSITEADDRESS(os.Args[1]))

	err := downloadfiles.DownloadFile(url, filename)
	if err != nil {
		panic(err)
	}

}

func main() {
	fmt.Println(geturldata.DownloadWEBADDRESS())
	// parseWEBSITEADDRESS()
	gettingAlbumTitles()
}
