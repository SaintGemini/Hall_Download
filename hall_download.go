package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getFile(vol string) {
	// Make HTTP request
	response, err := http.Get(vol)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var name string
	if strings.Contains(vol, "all-seeing-eye") {
		name = strings.TrimPrefix(vol, "http://www.manlyhall.org/all-seeing-eye/")
	} else if strings.Contains(vol, "horizon") {
		name = strings.TrimPrefix(vol, "http://www.manlyhall.org/horizon/")
	} else if strings.Contains(vol, "prs-journal") {
		name = strings.TrimPrefix(vol, "http://www.manlyhall.org/prs-journal/")
	} else {
		name = "output.pdf"
	}
	// Create output file
	outFile, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	// Copy data from HTTP response to file
	_, err = io.Copy(outFile, response.Body)
	if err != nil {
		log.Fatal(err)
	}
}

// This will get called for each HTML element found
func processElement(index int, element *goquery.Selection) {
	// See if the href attribute exists on the element
	href, exists := element.Attr("href")
	if exists && strings.HasSuffix(href, ".pdf") {
		getFile(href)
		fmt.Println("Successfully download " + href)
	}
}

func main() {
	// Make HTTP request
	response, err := http.Get("https://manlyphall.info/journals-index-opt.htm")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	// Find all links and process them with the function
	// defined earlier
	document.Find("a").Each(processElement)

}
