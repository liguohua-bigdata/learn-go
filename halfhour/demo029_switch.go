package main

import (
	"fmt"
	"runtime"
)
/*
你可能已经猜到 switch 可能的形式了。
case 体会自动终止，除非用 fallthrough 语句作为结尾。
 */

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}