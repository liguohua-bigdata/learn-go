package main

import "fmt"

func main() {
	fmt.Println("***************************")
	{
		//可以自动类型推断
		var v1 = 10;
		fmt.Println(v1)
	}

	fmt.Println("***************************")
	{
		//声明和赋值放到一起
		v2 := "zhangsan"
		fmt.Println(v2)
	}

	fmt.Println("***************************")
	{
		//多重赋值，交换i和j变量
		var i = 12
		var j = 22
		fmt.Println(i)
		fmt.Println(j)
		i, j = j, i;
		fmt.Println(i)
		fmt.Println(j)
	}

	fmt.Println("***************************")
	{
		//匿名变量
		_, _, nickName := GetName()
		fmt.Println(nickName)
	}

}

func GetName() (firstName, lastName, nickName string) {
	//多重返回值
	return "May", "Chan", "Chibi Maruko"
}
