package main

import "fmt"

func main() {
	arrFunc()
}

func arrFunc() {
	var arrA [3]int
	arrA[0] = 1
	arrA[1] = 2
	fmt.Println(arrA)

	arrB := [3]int{4, 5, 6}
	fmt.Println(arrB)

	for key, val := range arrB {
		fmt.Println(key, val)
	}

	for i := 0; i < len(arrB); i++ {
		fmt.Println(arrB[i])
	}
}
