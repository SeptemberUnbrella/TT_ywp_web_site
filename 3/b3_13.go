package main

import "fmt"

func showMemoryAddress(x *int) { //传输参数为指针，通*int定义类型，传承需要&取内存地址
	fmt.Println(x)
	return
}

func main() {
	var i int = 1
	fmt.Println(&i)
	showMemoryAddress(&i)
}
