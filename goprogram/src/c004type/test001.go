package main

import "fmt"

func main() {
	fmt.Println("*************************001")
	{
		var b = true//bool
		fmt.Println(b)
	}
	fmt.Println("*************************002")
	{
		var b = 1 > 2//bool
		fmt.Println(b)
	}
	fmt.Println("*************************003")
	{
		//布尔类型不能接受其他类型的赋值，不支持自动或强制的类型转换。
		/*
		var b bool
		b=1// 编译错误
		b = bool(1) // 编译错误
		*/

	}

}
