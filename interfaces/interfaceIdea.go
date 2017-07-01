package main

import "fmt"
import "math"

type Geometry interface {
	area() float64
	perimeter() float64
}

type Rectangle struct {
	length float64
	width  float64
}

type Circle struct {
	radius float64
}

func (rectangle Rectangle) area() float64 {
	return rectangle.length * rectangle.width
}

func (rectangle Rectangle) perimeter() float64 {
	return 2 * (rectangle.length + rectangle.width)
}

func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

func (circle Circle) perimeter() float64 {
	return 2 * math.Pi * circle.radius
}

func Print(obj Geometry) {
	fmt.Println(obj)
	fmt.Println(obj.area())
	fmt.Println(obj.perimeter())
}

func main() {
	r := Rectangle{length: 4, width: 3}
	c := Circle{radius: 5}

	Print(r)
	Print(c)
}
