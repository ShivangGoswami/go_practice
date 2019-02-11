package main

import (
	"fmt"
)

type triangle struct {
	base   float64
	height float64
}
type square struct {
	sideLength float64
}

func main() {
	t := triangle{base: 5, height: 5}
	s := square{sideLength: 5}
	printArea(t)
	printArea(s)

}

func (t triangle) getArea() float64 {
	return t.base * t.height * 0.5
}
func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

type shape interface {
	getArea() float64
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
