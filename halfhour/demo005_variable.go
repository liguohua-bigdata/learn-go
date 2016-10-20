package main

import "fmt"
//定义变量,不初始化值,有默认值
var x, y, z int
var c, python, java bool

//定义变量,并初始化值
var x0, y0, z0 int = 1, 2, 3
var c0, python0, java0 = true, false, "no!"
func main() {
	//打印变量
	fmt.Println(x, y, z, c, python, java)
	//打印变量
	fmt.Println(x0, y0, z0, c0, python0, java0)
}