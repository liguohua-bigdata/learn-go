package main
/*
常量的定义与变量类似，只不过使用 const 关键字。常量可以是字符、字符串、布尔或数字类型的值。
 */
import "fmt"
//定义常量
const Pi = 3.14

func main() {
	//定义常量
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")
	//定义常量
	const Truth = true
	fmt.Println("Go rules?", Truth)
}