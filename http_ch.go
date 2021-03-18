package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	URL  string
	Size int
}

func main() {
	channel := make(chan Page)
	urls := []string{"https://example.com", "http://google.com", "https://golang.org", "https://golang.org/doc"}
	for _, url := range urls {
		go responseSize(url, channel)
	}
	for i := 0; i < len(urls); i++ {
		page := <-channel
		fmt.Printf("%s: %d\n", page.URL, page.Size)
	}
}

func responseSize(url string, channel chan Page) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal()
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal()
	}
	channel <- Page{URL: url, Size: len(body)}
}
