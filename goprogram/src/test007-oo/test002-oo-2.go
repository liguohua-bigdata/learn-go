package main

import "fmt"

type Rect struct {
	x, y          float32
	width, length float64
}

func (r Rect)Area() float64 {
	return r.width * r.length
}
func main() {
	//始化Rect类型的对象实例
	rect1 := new(Rect)
	sum := rect1.Area()
	fmt.Println(sum)

	//始化Rect类型的对象实例
	rect2 := &Rect{}
	sum = rect2.Area()
	fmt.Println(sum)

	//始化Rect类型的对象实例
	rect3 := &Rect{0, 0, 100, 200}
	sum = rect3.Area()
	fmt.Println(sum)

	//始化Rect类型的对象实例
	rect4 := &Rect{width: 100, length: 200}
	sum = rect4.Area()
	fmt.Println(sum)

}
