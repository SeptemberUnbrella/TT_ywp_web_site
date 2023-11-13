package main

import (
	"fmt"
	"reflect" //一个模块
)

func main() {
	var s string = "string"
	var i int = 10
	var f float32 = 1.2

	fmt.Println(reflect.TypeOf(s)) //输出s的类型
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(f))
}
