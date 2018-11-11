package main

import (
	fmt "fmt"
	"os"
)

type shape interface {
	area() float32
	print()
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

func (c circle) print() {
	fmt.Printf("%#v\n", c)
}

func (r ractangle) area() float32 {
	return r.height * r.width
}

func (r ractangle) print() {
	fmt.Printf("%#v\n", r)
}

func calArea(s shape) float32 {
	return s.area()
}

func print(s shape) {
	s.print()
}

func main() {
	c := circle{10}
	r := ractangle{10, 20}
	fmt.Println(calArea(c))
	fmt.Println(calArea(r))
	print(c)
	print(r)

	fmt.Printf("%q\n", "\"string\"")
	fmt.Printf("%s\n", "\"string\"")
	fmt.Printf("%p\n", &c)

	fmt.Printf("|%6d|%6d|\n", 12, 345)
	fmt.Printf("%15s\n", "---------------")
	fmt.Printf("|%6.2f|%6.2f|\n", 12.0, 345.3)
	fmt.Printf("%15s\n", "---------------")
	fmt.Printf("|%-6.2f|%6.2f|\n", 12.0, 345.3)

	fmt.Fprintf(os.Stderr, "a %s\n", "error")
}
