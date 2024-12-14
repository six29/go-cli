package main

import (
	"fmt"
	"math"
)

type Circle struct {
	r float64
}

func Area(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

func CircleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

func main() {
	cr := Circle{r: 2}
	fmt.Printf("%.2f\n", CircleArea(cr))

	cr2 := Circle{r: 4}
	fmt.Printf("%.2f\n", CircleArea(cr2))

	fmt.Printf("%.2f\n", CircleArea(cr))

	c := Circle{r: 10}
	fmt.Println(Area(&c))

	c2 := &Circle{r: 3}
	fmt.Println(Area(c2))

	fmt.Println(Area(&c))
}
