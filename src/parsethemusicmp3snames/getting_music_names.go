package parsethemusicmp3snames

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// HTMLParser function takes in a html response and creates a htmlParser object.
func HTMLParser(response string) *html.Node {
	doc, err := html.Parse(strings.NewReader(response))
	if err != nil {
		fmt.Println("Can't parse the html5 utf-8 encoded response")
		os.Exit(1)
	}
	return doc

}

// var ReturnAnchorTags func(*html.Node)

//MainList Previous error was not to define MainList as a global Variable
//MainList exports all the mp3s in a list form
var MainList []string

//ReturnAnchorTags prints out the <a> tags from the parsed html document.
func ReturnAnchorTags(n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				// fmt.Println(a.Val)
				MainList = append(MainList, a.Val)
				// break
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ReturnAnchorTags(c)
		// fmt.Print(c)
	}

	return MainList
}

// MainListManipulator exports all the mp3s in a list form
// func MainListManipulator(n *html.Node) {
// 	for c := n.FirstChild; c != nil; c = c.NextSibling {
// 		ReturnAnchorTags(c)
// 		fmt.Print(c)
// 	}
// }
