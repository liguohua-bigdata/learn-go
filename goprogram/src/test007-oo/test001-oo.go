package main

import "fmt"

type MyInteger int
//面向过程的函数定义
func Add(a MyInteger, b MyInteger) MyInteger {
	return a + b
}
//面向对象的函数定义
func (a MyInteger) AddO(b MyInteger) MyInteger {
	return a + b
}
//面向对象的函数定义，改变对象的值,传递指针
func (a *MyInteger)AddToO(b MyInteger) (MyInteger) {
	//改变对象中的值
	(*a) = (*a) + b;
	return (*a)
}

func main() {
	{
		//面向过程的函数调用
		var sum = Add(3, 4)
		fmt.Println(sum)
	}
	{
		//面向对象的函数调用
		var a MyInteger = 3
		var sum = a.AddO(4)
		fmt.Println(sum)
	}
	{
		var  a  MyInteger=3
		fmt.Println(a)
		var sum  =a.AddToO(4)
		fmt.Println(sum)
		fmt.Println(a)

	}

}
