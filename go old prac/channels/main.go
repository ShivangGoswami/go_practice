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
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		//checkLink(link)
		go checkLink(link, c)
	}

	//	fmt.Println(<-c)
	//	fmt.Println(<-c)
	//	fmt.Println(<-c)
	//	fmt.Println(<-c)
	//	fmt.Println(<-c)

	//	fmt.Println(<-c)
	//for i := 0; i < len(links); i++ {
	//fmt.Println(<-c)
	for l := range c {
		//time.Sleep(time.Second)
		//go checkLink(l, c)
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	} //}
}

func checkLink(link string, c chan string) {
	//time loop
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		//	c <- "Might be down i think"
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	//c <- "yep it's up!"
	c <- link
}
