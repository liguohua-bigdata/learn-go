package main

import "fmt"

func main() {
	const Pi float64 = 3.14159265358979323846
	fmt.Println(Pi)

	const zero = 0.0
	fmt.Println(zero)

	const (
		size int64 = 1024
		eof = -1
	)
	fmt.Println(size)
	fmt.Println(eof)

	const u, v float32 = 0, 3
	fmt.Println(u)
	fmt.Println(v)

	const a, b, c = 3, 4, "foo" //a=3,b=4,c="foo", 无类型整型和字符串常量
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}
