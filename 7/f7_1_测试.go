package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name    string
	Age     int8
	Address string
}

func main() {
	Lilutong := Student{
		Name:    "李路通",
		Age:     12,
		Address: "重庆市开州区九龙山庄",
	}
	fmt.Println(Lilutong)
	fmt.Println(reflect.TypeOf(Lilutong))
}
