package main

import "fmt"

func main() {
	accumulate()

}

func accumulate() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func forformat() {
	j := 0
	for ; j < 10; j++ { // ++ 自增

		fmt.Println("for ...", j)
	}
}

func forRange() {
	str := "hi 你好 图灵大猫"
	for key, val := range str {
		fmt.Printf("%d %c \n", key, val)
	}
}
