package main

import "fmt"

func main() {
	defer fmt.Println("I am run after the function compreletes")
	fmt.Println("Hello world")
	fmt.Println("nihao,李路通")

}

//defer是一个很有用的Go语言功能，
//它能够让您在函数返回前执行另一个函数。函数在遇到return语句或到达函数末尾时返回。
//defer语句通常用于执行清理操作或确保操作（如网络调用）完成后再执行另一个函数。
