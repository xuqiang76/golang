package main

import "fmt"

func main() {
	forContinue()
}

func gotoJump() {

	i := 0
	sum := 0

jumpFlag:

	i++
	sum += i
	if i < 100 {
		goto jumpFlag
	}

	fmt.Println(sum)

}

func forBreak() {

loopStart:
	for i := 0; ; i++ {
		for j := 0; ; j++ {
			fmt.Println(i, j)

			if j >= 2 {
				break loopStart
			}
		}
	}

}

func forContinue() {

loopStart:
	for i := 0; i < 2; i++ {

		for j := 0; j < 100; j++ {
			fmt.Println(i, j)
			if j >= 2 {
				continue loopStart

			}
		}
	}
}
