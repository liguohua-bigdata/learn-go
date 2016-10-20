package main

import "fmt"

func main() {
	var a = 3
	if a < 5 {
	} else {
	}
	/*
	1. 条件语句不需要使用括号将条件包含起来();
	2. 无论语句体内有几条语句，花括号{}都是必须存在的;
	3. 左花括号{必须与if或者else处于同一行;
	4. 在if之后，条件语句之前，可以添加变量初始化语句，使用;间隔;
	5. 在有返回值的函数中，不允许将“最终的”return语句包含在if...else...结构中，否则会编译失败:
	 */
	/**
	编译失败的案例如下:
	 func example(x int) int {
		if x == 0 {
			return 5
		} else {
			return x
		}
	}


	 */

}
