package main

import (
	"fmt"
	"time"
)
/*
没有条件的 switch 同 switch true 一样。
这一构造使得可以用更清晰的形式来编写长的 if-then-else 链。
 */
func main() {
	t := time.Now()
	var h = t.Hour();
	fmt.Println(h)
	switch {

	case h < 12:
		fmt.Println("Good morning!")
	case h < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}