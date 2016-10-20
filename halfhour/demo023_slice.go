package main

import "fmt"

func main() {
	// len(a)=5
	a := make([]int, 5)
	printSlice("a", a)
	// len(b)=0, cap(b)=5
	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)

	//slice 可以通过“重新切片”来扩容（增长到容量上限）
	b = b[:cap(b)] // len(b)=5, cap(b)=5
	b = b[1:]      // len(b)=4, cap(b)=4
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}