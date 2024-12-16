package main

import "fmt"

var intPtr *int
var notPtr int
var extVal string = "Ext Val"
var ext *string

type config struct {
	ext string
}

func main() {

	ext := &extVal // assigns extVal address to ext
	c := config{
		ext: *ext, // assigns value of ext to the structs ext
	}
	fmt.Println("c", c.ext)

	intPtr := 10
	notPtr := 20
	// var x int
	x := &intPtr        // 10
	fmt.Println("x", x) // prints the address of intPtr
	fmt.Println(*x)     // Prints the value of whats in the address of intPtr

	fmt.Println("intPtr in main() -->", intPtr)                         // prints the value
	fmt.Println("&intPtr in main -->", &intPtr)                         // prints the address
	ptr(&intPtr)                                                        // pass the address
	fmt.Println("intPtr in main() after adding 1 in ptr() -->", intPtr) // 11

	fmt.Println(notPtr) // 20
	nPtr(notPtr)        // 21
	fmt.Println(notPtr) // still 20

}

func ptr(p *int) {
	fmt.Println("*p in ptr() -->", *p) // 10
	*p += 1
}

func nPtr(p int) {
	p += 1
	fmt.Println("p+1 in nPtr() -->", p) // 10
}
