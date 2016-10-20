package main

import "fmt"

func main() {

	fmt.Println("*************************001")
	{
		//声明数组
		var arr [3]int
		fmt.Println(arr)
		fmt.Println(len(arr))//数组长度
		fmt.Println(arr[0])//数组元素
		//C语言相同，数组下标从0开始，len(array)-1 则表示最后一个元素的下标。下面的示例遍历整型数组并逐个打印元素内容:
		for i := 0; i < len(arr); i++ {
			fmt.Println("1  Array element[", i, "]=", arr[i])

		}

		//Go语言还提供了一个关键字range，用于便捷地遍历容器中的元素。当然，数组也是range 的支持范围。
		for i, v := range arr {
			fmt.Println("2  Array element[", i, "]=", v)
		}

	}

}
