package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c{
		go func (link string)  {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(url string, c chan string) {
	time.Sleep(5 * time.Second)
	_, err := http.Get(url)

	if err != nil {
		fmt.Printf("can't connect " + url)
		c <- url
		return
	}

	fmt.Println("Succes connect " + url)
	c <- url
}
