package geturldata

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func returnURLBODY() {
	websiteAddress := os.Args[1]
	response, error := http.Get(string(websiteAddress))
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
