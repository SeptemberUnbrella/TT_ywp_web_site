package main

import "fmt"

func addition(x int, y int) int { //参数接收与返回均是int类型
	return x + y
}

func main() {
	var s string = "three"      //此处声明一个string类型的变量
	fmt.Println(addition(1, s)) //int类型与string类型不能想加
}
