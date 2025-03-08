package main

import (
	"fmt"
	"math"
)

// define our interface
type Shape interface {
	Area() float64
}

// implement the interface with a struct
type Rectangle struct {
	length, width float64
}

// implement area method in rectangle struct
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

type Circle struct {
	radius float64
}

// implement the area method in the circle struct
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// calling the Area func in our struct
func calculateArea(s Shape) float64 {
	return s.Area()
}

func main() {
	rect := Rectangle{width: 4, length: 7}
	cir := Circle{radius: 2}

	fmt.Println("calculate area:", calculateArea(rect))
	fmt.Println("calculate area:", calculateArea(cir))
}
