package main

import "fmt"

func main() {

	fmt.Println("*************************001")
	{
		//int和int32在Go语言里被认为是两种不同的类型，编译器也不会帮你自动 做类型转换

		var value2 int32
		value1 := 64 // value1将会被自动推导为int类型
		//value2 = value1 // 编译错误,cannot use value1 (type int) as type int32 in assignment
		value2 = int32(value1) // 编译通过
		fmt.Println(value2)

	}
	fmt.Println("*************************001")
	{
		//两个不同类型的整型数不能直接比较，比如int8类型的数和int类型的数不能直接比较
		{
			/*
			var i int32
			var j int64
			i, j = 1, 2
			if i==j {     //invalid operation: i == j (mismatched types int32 and int64)
			}
			*/

		}
		//各种类型的整型变量都可以直接与字面常量(literal)进行比较
		{
			var i int32
			var j int64
			i, j = 1, 2
			if i==1&&j==2 {
				fmt.Println("yes")
			}
		}

	}
	//取反在C语言中是~x，而在Go语言中 是^x。

}
