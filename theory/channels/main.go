package main

import "fmt"

func main() {

	urls := []string{
		"https://google.com",
		"https://amazon.com",
		"https://microsoft.com",
		"https://radhat.com",
	}

	// var results chan string = make(chan string)
	// results := make(chan string)
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			c <- url
		}(url)
	}
	for range urls {
		fmt.Println(<-c)
	}
}
