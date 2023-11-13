package main

import "fmt"

func addition(x int, y int) int {
	//此处声明接收两个x，y 的int类型的参数，返回值类型也为int
	return x + y
}

func main() {
	fmt.Println(addition(2, 4)) //此处调用参数传入2，4
}
