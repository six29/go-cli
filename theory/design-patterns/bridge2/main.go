// https://www.youtube.com/watch?v=tO1iZ3xbmNM

package main

import "fmt"

func main() {
	triangle := TriangularObject{
		height: 1,
		base:   2,
		length: 3,
	}
	rectangle := RectangularObject{
		x: 10,
		y: 12,
		z: 9,
	}
	x := 3
	var object Object
	if x > 2 {
		object = &triangle
	} else {
		object = &rectangle
	}

	volCalc := VolumeCalculator{
		object: object,
	}
	vol := volCalc.Volume()
	fmt.Println(vol)
}

// encapsulated object
type Object interface {
	Area() float64
	Height() float64
}

// feature object
type VolumeCalculator struct {
	object Object
}

func (calculator *VolumeCalculator) Volume() float64 {
	return calculator.object.Area() * calculator.object.Height()
}

// Triangular
type TriangularObject struct {
	height float64
	base   float64
	length float64
}

func (t *TriangularObject) Area() float64 {
	return t.height * t.base / 2
}

func (t *TriangularObject) Height() float64 {
	return t.length
}

// Rectangular object
type RectangularObject struct {
	x float64
	y float64
	z float64
}

func (r *RectangularObject) Area() float64 {
	return r.x * r.y
}

func (r *RectangularObject) Height() float64 {
	return r.z
}
