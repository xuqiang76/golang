package main

import "fmt"

var (
	i int
	j   = 1
	str = "text"
)

func varValue() {
	var m int
	var str string

	fmt.Println(m, str)
	fmt.Printf("%d %q \n", m, str)
	fmt.Println()
}

func varInit() {
	var m, n int = 1, 2
	var str string = "text"

	fmt.Println(m, n, str)
	fmt.Println()
}

func varNoType() {
	var m, n, s = 1, 2, "text"

	fmt.Println(m, n, s)
	fmt.Println()
}

func varShort() {
	m, n, s := 1, 2, "text"
	n = 5
	fmt.Println(m, n, s)
	fmt.Println()
}

func main() {

	varValue()
	varInit()
	varNoType()
	varShort()

}
