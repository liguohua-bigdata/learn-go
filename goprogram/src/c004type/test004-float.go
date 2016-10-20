package main

import (
	"fmt"
)

func main() {

	/*
	Go语言定义了两个类型float32和float64，
	其中float32等价于C语言的float类型， float64等价于C语言的double类型。
	 */
	fmt.Println("*************************001")
	{
		//1. 如果不加小数点，fvalue2会被推导为整型而不是浮点型
		//2.类型将被自动设为float64， 而不管赋给它的数字是否是用32位长度表示的
		fvalue2 := 12.0
		fmt.Println(fvalue2)

		var fvalue1 float32
		fvalue1 = 12
		//强制类型转换:
		fvalue1 = float32(fvalue2)
		fmt.Println(fvalue1)
	}
	fmt.Println("*************************001")
	{
		//浮点数不是一种精确的表达方式，所以像整型那样直接用==来判断两个浮点数是否相等 是不可行的，这可能会导致不稳定的结果。
		//一种推荐的替代方案:
		/*
		import "math"
		// p为用户自定义的比较精度，比如0.00001
		func IsEqual(f1, f2, p float64) bool {
		return math.Fdim(f1, f2) < p
		}
		*/

	}


}
