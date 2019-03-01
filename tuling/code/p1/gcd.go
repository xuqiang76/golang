package main

import "fmt"

//辗转相除法求最大公约数
func gcd(x, y int) int {
	var tmp int
	tmp = x % y
	if tmp > 0 {
		return gcd(y, tmp)

	} else {
		return y
	}
}

func main() {
	var x, y int
	x, y = 12, 15
	fmt.Println("gcd =", gcd(x, y))
}
