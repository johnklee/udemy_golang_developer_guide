package main

import "fmt"

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

type shape interface {
	getArea() float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return t.height * t.base * 0.5
}

func printArea(s shape) {
	fmt.Printf("The area of give shape is %f\n", s.getArea())
}

func main() {
	s := square{sideLength: 5}
	t := triangle{height: 3, base: 4}
	for _, aShape := range []shape{s, t} {
		printArea(aShape)
	}
}
