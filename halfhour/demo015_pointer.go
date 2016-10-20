package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	/*
	Go 有指针，但是没有指针运算。
	结构体字段可以通过结构体指针来访问。通过指针间接的访问是透明的。
	*/
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9

	//访问变量
	fmt.Println(v)
	//访问指针
	fmt.Println(p)
}