package main

import "fmt"

func main() {
	ifFunc()
}

func ifFunc() {
	a := 11
	if a > 10 {
		fmt.Println(">10")
	} else if a < 0 {
		fmt.Println("<0")
	} else {
		fmt.Println("[0,10]")
	}
}

func switchFunc() {
	a := 10
	switch a {
	case 10:
		fmt.Println("ten")
	case 9, 8, 7:
		fmt.Println("nine eight seven")
	default:
		fmt.Println("other")
	}

	switch {
	case a > 0 && a < 60:
		fmt.Println("F")
		//fallthrough
	case a < 80:
		fmt.Println("C")
	case a < 90:
		fmt.Println("B")
	case a <= 100:
		fmt.Println("A")
	default:
		fmt.Println("error")
	}
}
