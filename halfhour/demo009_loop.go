package main

import "fmt"
/**
Go 只有一种循环结构，for 循环。
基本的 for 循环看起来跟 C 或者 Java 中做的一样，除了没有了 ( ) 之外（甚至强制不能使用它们），而 { } 是必须的。
 */
func loop1() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
/**
跟 C 或者 Java 中一样，可以让前置、后置语句为空。
 */
func loop2() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}
/**
基于此可以省略分号： C 的 while 在 Go 中也是用 for 实现。
 */
func loop3() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
func loop4() {
	//如果省略了循环条件，它就是个死循环源。
	for ; ; {
		fmt.Println("dead loop4")
	}
	//而为了避免累赘，分号可以省略，因此一个死循环可以简洁地表达。
	for {
		fmt.Println("dead loop4")
	}
}
func main() {
	loop1()
	loop2()
	loop3()
	//loop4()
}