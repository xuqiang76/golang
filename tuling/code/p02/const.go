package main

import (
	"fmt"
)

func main() {
	enums()
}

const A = 1
const B = A + 2

func consts() {
	var i uint8 = 255
	fmt.Println(i+A, B)

}

const (
	_ = iota

	CAT = iota
	MOUSE
	DOG
)

func enums() {

	fmt.Println(CAT, MOUSE, DOG)
}
