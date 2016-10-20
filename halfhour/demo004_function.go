package main

import "fmt"
//定义函数,两个入参,返回一个整形
func add1(x int, y int) int {
	return x + y
}
//当两个或多个连续的函数命名参数是同一类型，则除了最后一个类型之外，其他都可以省略。
func add2(x, y int) int {
	return x + y
}
//函数可以返回任意数量的返回值。
func swap(x, y string) (string, string) {
	return y, x
}

//在 Go中，函数可以返回多个“结果参数”，而不仅仅是一个值。它们可以像变量那样命名和使用。
//如果命名了返回值参数，一个没有参数的 return 语句，会将当前的值作为返回值返回。
func split(sum int) (x, y int) {
	x = sum * 4/9
	y = sum - x
	return//不能省略
}
func main() {
	//调用add1
	fmt.Println(add1(42, 13))
	//调用add2
	fmt.Println(add2(42, 13))
	//调用swap
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	//调用split
	fmt.Println(split(17))
}