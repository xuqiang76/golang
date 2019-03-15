package main

import (
	"fmt"
)

func consts() {
	const a, b, c = 3, true, "hello"
	const i int = 4

	fmt.Println(a,b,c)
	fmt.Println(i)

}

func menu(){
	const (
		_ = iota
		man = iota
		women
		child
		dog
	)
	fmt.Println(man, women, child,dog)
}

func main() {
	consts()

	menu()

}
