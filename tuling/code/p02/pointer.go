package main

import "fmt"

func main() {
	varPointer()
}


func varPointer() {
	a := 2
	var ptr *int = &a
	b := *ptr
	*ptr = 3

	fmt.Println(a, b, ptr)
}

func newPointer() {
	ptr := new(string)
	*ptr = "text"

	fmt.Println(*ptr, ptr)
}
