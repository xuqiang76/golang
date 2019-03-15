package main

import (
	"fmt"
)

func varComplex() {
	a := 3 + 4i
	b := complex(3, 4)
	fmt.Println(a)
	fmt.Println(b)
}

func varZero() {
	var (
		a int
		b float64
		c complex128
		d bool
		e string
	)

	fmt.Println(a, b, c, d, e)
	fmt.Printf("%q\n", e)
}

func varTypeChange() {
	a := 1.23
	b := int(a)
	fmt.Println(b)
}

func swap() {

	x, y := 1, 2

	tmp := x
	x = y
	y = tmp

	//x = x ^ y
	//y = x ^ y
	//x = x ^ y


	fmt.Println(x, y)
}

func main() {
	varComplex()
	varZero()
	varTypeChange()
	swap()
}
