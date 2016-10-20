package main

import "fmt"

func main() {
	{
		//test(0)//0
		//test(1)//1
		//test(2)//3
		//test(4)//4, 5, 6
		//test(9)//Default
	}
	{
		test2(1);//t0-3
	}
	{
		/**
		左花括号{必须与switch处于同一行;
		条件表达式不限制为常量或者整数;
		单个case中，可以出现多个结果选项;
		与C语言等规则相反，Go语言不需要用break来明确退出一个case;
		只有在case中明确添加fallthrough关键字，才会继续执行紧跟的下一个case; 可以不设定switch之后的条件表达式，在此种情况下，整个switch结构与多个
		if...else...的逻辑作用等同。
		 */
	}

}

func test2(Num int) {
	//switch后面的表达式甚至不是必需的
	switch {
	case 0 <= Num && Num <= 3:
		fmt.Printf("0-3")
	case 4 <= Num && Num <= 6:
		fmt.Printf("4-6")
	case 7 <= Num && Num <= 9:
		fmt.Printf("7-9")
	}
}
func test(i int) {

	switch i {
	case 0:
		fmt.Printf("0")
	case 1:
		fmt.Printf("1")
	case 2:
		fallthrough
	case 3:
		fmt.Printf("3")
	case 4, 5, 6:
		fmt.Printf("4, 5, 6")
	default:
		fmt.Printf("Default")
	}
}
