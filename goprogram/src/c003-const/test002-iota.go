package main

import "fmt"

func main() {
	fmt.Println("*************************001")
	{
		const (
			// iota被重设为0
			c0 = iota //0
			c1 = iota //1
			c2 = iota //2
		)
		fmt.Println(c0)
		fmt.Println(c1)
		fmt.Println(c2)
	}

	fmt.Println("*************************002")
	{
		const (
			//iota在每个const开头被重设为0
			a = 1 << iota //iota=0
			b = 1 << iota //iota=1
			c = 1 << iota //iota=2
		)
		fmt.Println(a)//1
		fmt.Println(b)//2
		fmt.Println(c)//4

	}
	fmt.Println("*************************003")
	{
		//iota在每个const开头被重设为0
		const x = iota //iota=0
		const y = iota //iota-0
		fmt.Println(x)
		fmt.Println(y)

	}
	fmt.Println("*************************004")
	{
		//如果两个const的赋值语句的表达式是一样的，那么可以省略后一个赋值表达式。
		const (
			// iota被重设为0
			c0 = iota //0
			c1  //1
			c2  //2
		)
		fmt.Println(c0)
		fmt.Println(c1)
		fmt.Println(c2)

	}
	fmt.Println("*************************005")
	{
		//如果两个const的赋值语句的表达式是一样的，那么可以省略后一个赋值表达式。
		const (
			//iota在每个const开头被重设为0
			a = 1 << iota //iota=0
			b //iota=1
			c //iota=2
		)
		fmt.Println(a)//1
		fmt.Println(b)//2
		fmt.Println(c)//4
	}
	fmt.Println("*************************001")
	{
		//形成枚举，Go语言并不支持众多其他语言明确支持的enum关键字。
		const (
			Sunday = iota
			Monday
			Tuesday
			Wednesday
			Thursday
			Friday
			Saturday
			numberOfDays
		)
		fmt.Println(Sunday)//0
		fmt.Println(Monday)//1
		fmt.Println(Tuesday)//2
		fmt.Println(Wednesday)//3
		fmt.Println(Thursday)//4
		fmt.Println(Friday)//5
		fmt.Println(Saturday)//6
		fmt.Println(numberOfDays)//7
	}
	fmt.Println("*************************001")
	{

	}
}
