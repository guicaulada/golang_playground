package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	// sequential check
	// for _, link := range links {
	// 	checkLink(link)
	// }

	// unwaited goroutine
	// for _, link := range links {
	// 	go checkLink(link)
	// }

	// goroutine with channel
	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}

	// print up to 5 received values
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	// infinite loop
	// for {
	// 	go checkLink(<-c, c)
	// }

	// infinite loop
	// runs every time c receives a value
	for l := range c {
		// go checkLink(l, c)
		// using a function literal to include a delay
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c) // receive l as an argument
			// checkLink(l, c) / /capture l from the other goroutine
			// we can't do this because the value of l changes
			// we shouldn't share variables between goroutines
			// if we want to do that we should use a channel
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}
