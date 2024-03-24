package main

import "fmt"

type shape interface {
	getArea() float64
}
type triangle struct {
	base   float64
	height float64
}

type square struct {
	sideLength float64
}

func main() {

	tr := triangle{10, 10}
	sq := square{20}

	printArea(tr)
	printArea(sq)

}

func printArea(s shape) {
	fmt.Printf("%.2f\n", s.getArea())
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.height * t.base
}
