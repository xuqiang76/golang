package main

import (
	"fmt"
)

func consts() {
	const a, b, c = 3, true, "hello"
	const i int = 4

	fmt.Println(a, b, c)
	fmt.Println(i)
	fmt.Printf("%T %T \n", a, b)

}

func menu() {
	const (
		cat   = iota
		man   = 2.3
		women = iota
		_
		child
		dog
	)

	fmt.Println(cat, man, women, child, dog)
}

func main() {
	consts()

	menu()

}
