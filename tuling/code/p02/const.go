package main

import (
	"fmt"
)

func main() {
	consts()
	enums()
}

const A  = 2

func consts() {

	var i uint8 = 255
	fmt.Println(i + A)

}

const (
	CAT = iota
	_

	MOUSE
	DOG
)

func enums() {

	fmt.Println(CAT, MOUSE, DOG)
}


