package main

import (
	"fmt"
	"sort"
)

//1.声明Hero结构体
type Hero struct {
	Name string
	Age  int
}

//2.声明一个hero结构体切片类型
type HeroSlice []Hero

func main() {
	//先定义一个数组
	var intSlice = []int{0, -1, 10, 7, 90}
	sort.Ints(intSlice)
	fmt.Println(intSlice)
}
