package main

import (
	"fmt"
)

func main() {

	{
		sum, e := Add(3, 4)
		fmt.Println(sum, e)
	}

	{
		sum2, e2 := Add2(3, 4, 5, 6, 7, 8, 9)
		fmt.Println(sum2, e2)
	}
	{
		arr := []int{3, 4, 5, 6, 7, 8, 9};
		//传递数组要打散
		sum2, e2 := Add2(arr...)
		fmt.Println(sum2, e2)

	}
	{
		var v1 int = 1
		var v2 int64 = 234
		var v3 string = "hello"
		var v4 float32 = 1.234
		var v5 bool = true
		MyPrint(v1, v2, v3, v4, v5)
	}
	{
		//匿名函数由一个不带函数名的函数 明和函数体组成，
		//func(a, b int, z float64) bool {
		//	return a * b < int(z)
		//}
	}
	{
		//立即执行的匿名函数
		var b bool  = func(a, b int, z float64) bool {
			return a * b < int(z)
		}(3,5,12)//立即执行
		fmt.Println(b)
	}
	{
		//匿名函数可以直接赋值给一个变量或者直接执行:]
		f := func(x, y int) int { return x + y
		}
		fmt.Println(f(3,5))
	}


}
func MyPrint(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case int32:
			fmt.Println(arg, "is an int32 value.")
		case int64:
			fmt.Println(arg, "is an int64 value.")
		case bool:
			fmt.Println(arg, "is an bool value.")
		default:
			fmt.Println(arg, "is an float32 value.")

		}
	}

}


//返回值的参数名称可以去掉，...type本质上是一个数组切 ，也就是[]type，只能作为最后一个参数
func Add2(args ...int) (int, error) {
	sum := 0
	for _, arg := range args {
		sum += arg
	}
	return sum, nil
}
//多重返回值
func Add(a int, b int) (sum int, e error) {
	return a + b, nil
}
