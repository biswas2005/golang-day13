package basics

import (
	"fmt"
	"math"
)

type shape interface {
	Area() float64
}

type Rectangle struct {
	height, width float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func (r Circle) Area() float64 {
	return math.Pi * r.radius * r.radius
}

func ShapeArea() {
	r := Rectangle{height: 5, width: 8}
	c := Circle{radius: 10}

	Shapes := []shape{r, c}

	for _, val := range Shapes {
		fmt.Printf("Area of %T: %.2f\n", val, val.Area())
	}
}
