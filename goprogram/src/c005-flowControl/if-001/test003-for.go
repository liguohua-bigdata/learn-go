package main

import "fmt"

func main() {
	/**
	Go语言中的循环语句只支持for关键字，而不支持while和do-while 结构
	 */
	fmt.Println("*************************001")
	{
		sum := 0
		for i := 0; i <= 100; i++ {
			sum += i
		}
		fmt.Println(sum)
	}
	fmt.Println("*************************002")
	{
		//无限循环的场景
		sum := 0
		i := 0
		for {
			sum += i
			i++
			if i > 100 {
				break;
			}
		}
		fmt.Println(sum)
	}
	fmt.Println("*************************001")
	{
		//条件表达式中也支持多重赋值
		a := []int{31, 72, 13, 48, 50, 68}
		fmt.Println("befor=", a)
		for i, j := 0, len(a) - 1; i < j; i, j = i + 1, j - 1 {
			a[i], a[j] = a[j], a[i]
		}
		fmt.Println("after=", a)

	}
	/**
	1. 左花括号{必须与for处于同一行。
	2. Go语言中的for循环与C语言一样，都允许在循环条件中定义和初始化变量，唯一的区别
	   是，Go语言不支持以逗号为间隔的多个赋值语句，必须使用平行赋值的方式来初始化个变量。
	 */
	fmt.Println("*************************003")
	{
		/**
		Go语言的for循环同样支持continue和break来控制循环，但是它提供了一个
		更高级break，可以选择中断哪一个循环，
		 */
		var sum = 100;
		JLoop:
		for j := 0; j < 5; j++ {
			for i := 0; i < 10; i++ {
				sum = i * j
				if sum > 50 {
					break JLoop
				}
			}
		}
		sum = sum * sum
		fmt.Println(sum)

	}
	fmt.Println("*************************004")
	{
		/**
		goto语句被多数语言学者所反对，谆谆告诫不要使用。但对于Go语言这样一个惜关键字如 6 金的语言来说，居然仍然支持goto关键字，
		无疑让某些人跌破眼镜。但就个人一年多来的Go语 言编程经验来说，goto还是会在一些场合下被证明是最合适的。
		 */
		i := 0
		HERE:
		fmt.Println(i)
		i++
		if i < 10 {
			goto HERE
		}
	}
}
