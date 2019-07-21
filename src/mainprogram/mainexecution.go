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

func gettingAlbumTitles() []string {
	return parsethemusicmp3snames.ReturnAnchorTags(parsethemusicmp3snames.HTMLParser(parseWEBSITEADDRESS()))

}

func listManipulator() []string {
	var FinalList []string
	lengthOfMainList := len(gettingAlbumTitles())
	for i := 0; i < lengthOfMainList; i++ {
		FinalList = append(FinalList, geturldata.DownloadWEBADDRESS())
	}
	for i := 0; i < lengthOfMainList; i++ {
		FinalList[i] = FinalList[i] + gettingAlbumTitles()[i]
	}

	return FinalList
}

func downloadTheWholeAlbum() {
	changeDirectory()
	var finallist []string
	finallist = listManipulator()
	for i := 1; i < len(finallist); i++ {
		// fmt.Print(len(listManipulator()))
		fmt.Printf("####################################################################################\n")
		fmt.Printf("DOWNLOADING THE %d SONG\n", i)
		fmt.Printf("%s\n", string(gettingAlbumTitles()[i]))
		downloaderASSEMBLY(finallist[i], i)

	}
}
func downloaderASSEMBLY(url string, i int) {
	if len(os.Args) != 3 {
		fmt.Println("usage: ./mainexecution therockmusic.orgAlbumURL foldername")
		os.Exit(1)
	}
	// url := os.Args[1]
	filename := string(gettingAlbumTitles()[i])

	err := downloadfiles.DownloadFile(url, filename)
	if err != nil {
		panic(err)
	}

}

// Changing filename to filepath so that the songs are downloaded
// in separate folder Songs and then in the albums folder .
func changeDirectory() {
	os.Mkdir(geturldata.ParseWEBSITEADDRESS(os.Args[1]), 777)
	os.Chdir("./" + geturldata.ParseWEBSITEADDRESS(os.Args[1]))
}

func makesongsdirectory() {
	os.Chdir("../")
	os.Mkdir("songs", 777)
	os.Chdir("./songs")

}

func main() {
	// fmt.Println(geturldata.DownloadWEBADDRESS())
	// parseWEBSITEADDRESS()
	//fmt.Print(gettingAlbumTitles())
	//fmt.Print(listManipulator())
	makesongsdirectory()
	downloadTheWholeAlbum()
}
