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

/*empty interfaces in Go (works like ANY Type)
Go's empty interface (interface{}) is a wildcard that accepts any type.
*/

func printValue(value interface{}) {
	fmt.Println(value)
}

func main() {
	rect := Rectangle{width: 4, length: 7}
	cir := Circle{radius: 2}

	fmt.Println("calculate area:", calculateArea(rect))
	fmt.Println("calculate area:", calculateArea(cir))

	printValue(40)
	printValue("ichami")
	printValue(true)

	printType("ichami")

	descripeShape(rect)
}

/*Type Assertion (extracting origin type from interface)
When using interface{}, we often need to extract the original type.
*/

func printType(i interface{}) {
	value, ok := i.(string) //checks if the type is string
	if ok {
		fmt.Println("string value:", value)
	} else {
		fmt.Println("not a string")
	}
}

// interface embedding
type Measurable interface {
	Perimeter() float64
}

type Geometry interface {
	Shape
	Measurable
}

func (r Rectangle) Perimeter() float64 {
	return r.length + r.width
}

func descripeShape(g Geometry) {
	fmt.Println("Area:", g.Area())
	fmt.Println("Perimeter:", g.Perimeter())
}
