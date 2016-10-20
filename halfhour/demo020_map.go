package main

import "fmt"

func main() {
	/**
	如果 key 在 m 中， ok 是 true。 否则，ok 是 false 并且 elem 是 map 的元素类型的零值。
	同样的，当从 map 中读取某个不存在的键时，结果是 map 的元素类型的零值。
	 */
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}