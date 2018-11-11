package main

import (
	"fmt"
	"mymath"
)

type A struct {
	name string
	age  int
}

func (a A) toString() {
	fmt.Printf("name %v - age %v", a.name, a.age)
}

func sum(a int, b int) int {
	return a + b
}

func solvePtbh() {
	var a, b, c int
	fmt.Print("Enter a")
	_, err := fmt.Scanf("%d", &a)
	fmt.Print("Enter b")
	_, err = fmt.Scanf("%d", &b)
	fmt.Print("Enter c")
	_, err = fmt.Scanf("%d", &c)

	if err != nil {
		fmt.Print(err)
	}

	delta := float64(b*b - 4*a*c)
	if delta > 0 {
		x1 := (float64(-b) + mymath.Sqrt(delta)) / float64(2*a)
		x2 := (float64(-b) - mymath.Sqrt(delta)) / float64(2*a)
		fmt.Printf("x1 = %v, x2 = %v\n", x1, x2)
	} else if delta == 0 {
		x := -b / (2 * a)
		fmt.Printf("nghiem kep x = %v\n", x)
	} else {
		fmt.Print("Phuong trinh vo nghiem")
	}
}

func main() {
	a := A{"tri", 24}
	a.toString()
}
