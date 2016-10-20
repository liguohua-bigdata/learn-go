package main

import "fmt"

func main() {
	fmt.Println("*************************001")
	{
		// 先定义一个数组
		var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		// 基于数组创建一个数组切片
		var mySlice []int = myArray[:5]

		//print myArray
		fmt.Println("Elements of myArray: ")
		for _, v := range myArray {
			fmt.Print(v, " ")
		}
		//print mySlice
		fmt.Println("\nElements of mySlice: ")
		for _, v := range mySlice {
			fmt.Print(v, " ")
		}
		fmt.Println()

		//Go语言支持用myArray[first:last]这样的方式来基于数组生成一 个数组切片
		mySlice = myArray[:]//基于myArray的所有元素创建数组切片:
		fmt.Println(mySlice)

		mySlice = myArray[:5]//[0,5)
		fmt.Println(mySlice)

		mySlice = myArray[5:]//从第5个元素开始的所有元素创建数组切片:
		fmt.Println(mySlice)

		mySlice = myArray[4:8]//从第4到第8
		fmt.Println(mySlice)

	}
	fmt.Println("*************************002")
	{
		/**
		并非一定要事先准备一个数组才能创建数组切片。Go语言提供的内置函数make()可以用于
		灵活地创建数组切片。
		 */
		//创建一个初始元素个数为5的数组切片，元素初始值为0:
		mySlice1 := make([]int, 5)
		fmt.Println(mySlice1)
		//创建一个初始元素个数为5的数组切片，元素初始值为0，并预留10个元素的存储空间:
		mySlice2 := make([]int, 5, 10)
		fmt.Println(mySlice2)
		fmt.Println("len=", len(mySlice2))//len()函数返回的是数组切片存储的元素个数。
		fmt.Println("cap=", cap(mySlice2))//cap()函数返回的是数组切片分配的空间大小。

		//直接创建并初始化包含5个元素的数组切片:
		mySlice3 := []int{1, 2, 3, 4, 5}
		fmt.Println(mySlice3)

	}
	fmt.Println("*************************003")
	{
		//切片遍历,传统
		mySlice := []int{1, 2, 3, 4, 5}
		for i := 0; i < len(mySlice); i++ {
			fmt.Println("mySlice[", i, "] =", mySlice[i])
		}
	}
	fmt.Println("*************************004")
	{
		//切片遍历,range
		mySlice := []int{19, 2, 39, 4, 5}
		for i, v := range mySlice {
			fmt.Println("mySlice[", i, "] =", v)
		}

	}
	fmt.Println("*************************005")
	{
		mySlice := make([]int, 5, 10)
		fmt.Println("len(mySlice):", len(mySlice))
		fmt.Println("cap(mySlice):", cap(mySlice))
		//append element
		mySlice = append(mySlice, 1)
		fmt.Println("append element=", mySlice)
		//append elements
		mySlice = append(mySlice, 1, 2, 3)
		fmt.Println("append elements=", mySlice)
		//append array
		mySlice2 := []int{8, 9, 10}
		mySlice = append(mySlice, mySlice2...)
		fmt.Println("append array=", mySlice)

	}
	fmt.Println("*************************006")
	{
		//数组切片可以基于一个数组创建，数组切片也可以基于另一个数组切片创建
		oldSlice := []int{1, 2, 3, 4, 5}
		newSlice := oldSlice[:3] // 基于oldSlice的前3个元素构建新数组切片
		fmt.Println("newSlice=", newSlice)

		/**
		选择的oldSlicef元素范围甚至可以超过所包含的元素个数，比如newSlice
		可以基于oldSlice的前6个元素创建，虽然oldSlice只包含5个元素。只要这个选择的范围不超
	   	过oldSlice存储能力(即cap()返回的值)，那么这个创建程序就是合法的。
	   	newSlice中超出 oldSlice元素的部分都会填上0。
		 */

		newSlice = oldSlice[:cap(oldSlice)]
		fmt.Println("newSlice cap =", newSlice)

	}
	fmt.Println("*************************007")
	{
		/**
		数组切片支持Go语言的另一个内置函数copy()，用于将内容从一个数组切片复制到另一个 数组切片。
		如果加入的两个数组切片不一样大，就会按其中较小的那个数组切片的元素个数进行 复制。
		 */
		{
			slice1 := []int{1, 2, 3, 4, 5}
			slice2 := []int{5, 4, 3}
			copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
			fmt.Println("slice1=", slice1, "slice2", slice2)
		}
		{
			slice1 := []int{1, 2, 3, 4, 5}
			slice2 := []int{5, 4, 3}
			copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
			fmt.Println("slice1=", slice1, "slice2", slice2)
		}
	}
}
