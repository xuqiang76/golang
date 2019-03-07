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
		v12 = "text中文"

	)

	//复数
	var v31 complex64 = complex(12, 13.5)
	v32 := 15 + 0.1i

	fmt.Println("v21 is", v21)
	fmt.Println("v22 is", v22)
	fmt.Println("v23 is", v23)
	fmt.Println("v24 is", v24)
	fmt.Println("v25 is", v25)

	fmt.Println()
	fmt.Println("v11 is", v11)
	fmt.Println("v12 is", v12)

	fmt.Println()
	fmt.Println("v31 is", v31)
	fmt.Println("v32 is", v32)

	//零值
	fmt.Println()
	var a int32
	var b float32
	var c bool
	var d complex64
	var e string
	var f *int
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(f)
	fmt.Printf("%q",e)

}


