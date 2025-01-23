package main

import "fmt"

type square struct {
	hight float64
	width float64
}

type triangle struct {
	hight float64
	width float64
}

type shape interface {
	getArea() float64
}

func (t triangle) getArea() float64 {
	return (t.hight * t.width) / 2
}

func (s square) getArea() float64 {
	return (s.hight * s.width)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	s := square{10, 10}

	t := triangle{10, 10}

	printArea(s)
	printArea(t)
}
