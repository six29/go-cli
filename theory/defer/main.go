package main

import "fmt"

func deferMe(s string) string {
	return s
}

func main() {
	defer fmt.Println("Deferred 1")
	defer fmt.Println(deferMe("Deferred"))
	fmt.Println("before main ends")
	defer fmt.Println("Deferred 2")
	defer fmt.Println("Deferred 3")
	fmt.Println("main ends")
}
