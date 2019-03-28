package main

import "fmt"

func main() {
	newPointer()
}

func varPointer() {
	a := 2
	var ptr = &a
	b := *ptr
	*ptr = 3

	fmt.Println(a, b, ptr, &b)

}


func newPointer() {
	ptr := new(string)
	*ptr = "text"

	fmt.Println(*ptr, ptr)
}
