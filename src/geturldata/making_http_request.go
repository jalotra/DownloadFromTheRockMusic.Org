package geturldata

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// I am getting this /lib/media/music/http://therockmusic.org/albums/awakened-by-hope/index.html

// ReturnURLBODY Prints what is present in the returned RootUrl WEbpage.
func ReturnURLBODY(urltofinditscontents string) {
	response, error := http.Get(string(urltofinditscontents))
	if error != nil {
		log.Fatalln(error)
	}
	defer response.Body.Close()

	body, error := ioutil.ReadAll(response.Body)
	if error != nil {
		log.Fatalln(error)
	}
	log.Println(string(body))
}

// ParseWEBSITEADDRESS takes in websiteAddress and return the album name
// example : blackwave or promises
func ParseWEBSITEADDRESS(websiteAddress string) string {
	var AlbumTitle string
	AlbumTitle = websiteAddress
	// http://therockmusic.org/albums/blackwave/index.html
	AlbumTitle = strings.TrimPrefix(AlbumTitle, "http://therockmusic.org/albums/")
	AlbumTitle = strings.TrimSuffix(AlbumTitle, "/index.html")

	//NOw AlbumTitle should return only
	// BlackWave

	return AlbumTitle
}

// DownloadWEBADDRESS outputs the formatted string like this one
// example : http://therockmusic.org/lib/media/music/dwelling/
func DownloadWEBADDRESS() string {
	var rootURL string
	rootURL = "http://therockmusic.org/lib/media/music/"
	rootURL = rootURL + ParseWEBSITEADDRESS(os.Args[1])
	//rootURL now becomes example :: http://therockmusic.org/lib/media/music/dwelling/

	return rootURL

}
