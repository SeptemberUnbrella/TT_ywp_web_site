package main

import "fmt"

func getPrize() (int, string) { //此处定义了函数返回值是两个值，并定义其中的类型
	var (
		i int    = 32
		s string = "hello sansan"
	)
	return i, s

}

func main() {
	a, b := getPrize() //此处需要拿出两个变量来接函数返回的两个值
	fmt.Println(a, b)

}
