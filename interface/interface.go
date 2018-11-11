package main

import fmt "fmt"

type shape interface {
	area() float32
}

type ractangle struct {
	width, height float32
}

type circle struct {
	radius float32
}

func (c circle) area() float32 {
	return 3.14 * c.radius * c.radius
}

func (r ractangle) area() float32 {
	return r.height * r.width
}

func calArea(s shape) float32 {
	return s.area()
}

func main() {
	fmt.Println(calArea(circle{10}))
	fmt.Println(calArea(ractangle{10, 20}))
}
