package main

import "fmt"

func main() {

	//变量的定义
	var v21 int32
	var v22 int = 2
	var v23 = 3
	v24 := 4.6
	var v25 float32 = float32(v24)
	var (
		v11 = 3
		v12 = "text"
		v13 = true
	)

	//复数
	var v31 complex64 = complex(12, 13.5)
	v32 := 15 + 0.1i

	fmt.Println("v21 is", v21)
	fmt.Println("v22 is", v22)
	fmt.Println("v23 is", v23)
	fmt.Println("v24 is", v24)
	fmt.Println("v25 is", v25)

	fmt.Println("v11 is", v11)
	fmt.Println("v12 is", v12)
	fmt.Println("v13 is", v13)

	fmt.Println("v31 is", v31)
	fmt.Println("v32 is", v32)

	fmt.Println(swap(v22, v23))

}

func swap(x, y int) (int, int) {
	var a int
	a = x
	x = y
	y = a

	return x, y
}
