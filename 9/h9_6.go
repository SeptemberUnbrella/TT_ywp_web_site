package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 1
	var s string = "egg"
	var breakfast string = strconv.Itoa(i) + s //通过strconv中Itoa方法将其转换为字符串
	fmt.Println(breakfast)
}
