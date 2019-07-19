package main

import (
	"downloadfiles"
	"fmt"
	"geturldata"
	"os"
)

func parseWEBSITEADDRESS() {
	geturldata.ReturnURLBODY(geturldata.DownloadWEBADDRESS())
}

func downloaderASSEMBLY() {
	if len(os.Args) != 3 {
		fmt.Println("usage: download url filename")
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
	// fmt.Println(geturldata.DownloadWEBADDRESS())
	parseWEBSITEADDRESS()
}
