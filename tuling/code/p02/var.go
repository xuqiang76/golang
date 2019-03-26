package main

import "fmt"

func main() {

	varDeclare()

}

func varDeclare() {
	var m int
	var str string

	m = 3
	str = "text"

	fmt.Println(m, str)
}

func varInit() {
	var m, n int = 1, 2
	var str string = "text"
	fmt.Println(m, n, str)
}

func varShort() {
	m, n, s := 1, 2, 4
	n = 1
	a, n := 1, 666 // n 只做赋值(注意类型) ,不做声明

	fmt.Println(m, n, s, a)

}

var (
	i   int
	j   = 1
	str = "text1"
)
