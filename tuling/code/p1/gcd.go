package main

import "fmt"

//辗转相除法求最大公约数
func gcd(x, y uint) uint {
	var tmp uint
	tmp = x % y
	if tmp > 0 {
		return gcd(y, tmp)

	} else {
		return y
	}
}

func main() {
	var x, y uint
	x, y = 12, 15
	fmt.Println("gcd =", gcd(x, y))
	fmt.Println("gcd =", gcd(100, 120))
}
