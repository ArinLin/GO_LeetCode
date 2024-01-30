package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Square struct {
	X float64
}

func (s Square) Area() float64 {
	return s.X * s.X
}

// написать реализацию интерфейса для круга
type Circle struct {
	R float64
}

func (c Circle) Area() float64 {
	return math.Pi * (c.R * c.R)
}

func Calculate(x Shape) {
	fmt.Printf("Area: %f\n", x.Area())
}

func main() {
	// вывести расчет площадей для квадрата и для круга используя интерфейс
	var shape Shape
	shape = Circle{R: 64}
	Calculate(shape)
	shape = Square{X: 58}
	Calculate(shape)
}
