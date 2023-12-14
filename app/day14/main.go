package main

import "fmt"


type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width float64
	height float64
}

type Circle struct {
	radius float64
}



func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width * r.height)
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.radius
}

func PrintAreaAndPerimeter(s Shape) {
	fmt.Println(s.Area())
	fmt.Println(s.Perimeter())
}



func main() {
	var r Shape
    var c Shape
	r = Rectangle{width: 10, height: 20}
	c = Circle{radius: 10}
    PrintAreaAndPerimeter(r)
	PrintAreaAndPerimeter(c)
}

