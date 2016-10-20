package main

import "fmt"

func main() {

	fmt.Println("*************************001")
	{
		var str string // 声明一个字符串变量
		str = "Hello world" // 字符串赋值
		ch := str[0] // 取字符串的第一个字符
		fmt.Printf("The length of \"%s\" is %d \n", str, len(str))
		fmt.Printf("The first character of \"%s\" is %c.\n", str, ch)

		//字符串的内容不能在初始 化后被修改
		//str[0] = 'X' // 编译错误
		//但可以改变字符串的指向
		str = "Hello world2"
		fmt.Println(str)

	}
	fmt.Println("*************************002")
	{
		//字符串也支持声明时进行初始化的做法
		str := "good go"
		//取长度
		fmt.Println(len(str))
		//取首字母
		fmt.Println(str[0])

	}
	fmt.Println("*************************001")
	{
		{
			str := "good go"
			//字节数组的方式遍历:
			for i := 0; i < len(str); i++ {
				fmt.Println(str[i])//类型为byte
			}
		}
		{
			//Unicode字符遍历:
			str := "Hello,世界"
			for i, ch := range str {
				fmt.Println(i, ch)//ch的类型为rune
			}
		}

	}
	{
		/**
		在Go语言中支持两个字符类型，
		一个是byte(实际上是uint8的别名)，代表UTF-8字符串的单个字节的值;
		另一个是rune，代表单个Unicode字符。
		 */
	}
}
