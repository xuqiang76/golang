package main

import (
	"fmt"
)

func main() {

	varTypeChange()
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
	b := int(a) //bool值无法参与数值运算，也无法与其他类型进行转换
	fmt.Println(b)
}
