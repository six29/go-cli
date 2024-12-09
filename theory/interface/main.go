package main

import "fmt"

type Shape interface {
	Area() float32
}

type Rectangle struct {
	length float32
	width  float32
}

type Circle struct {
	radius float32
}

func (r *Rectangle) Area() float32 {
	return r.length * r.width
}

func (c *Circle) Area() float32 {
	return 2 * 3.14 * c.radius
}

func printArea(s Shape) {
	fmt.Println(s.Area())
}

func main() {
	rect := &Rectangle{10.5, 3.0}
	circ := &Circle{4.0}
	printArea(rect)
	printArea(circ)

}
