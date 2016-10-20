package main

import "fmt"
//函数是完全闭包的。
func adder() func(int) int {
	//函数 adder 返回一个闭包。每个闭包被绑定到了特定的 sum 变量上。
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-i),
		)
	}
}