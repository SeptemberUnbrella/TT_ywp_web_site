package main

import (
	"fmt"
	"reflect"
	"strconv" //用于字符串转换的模块
)

func main() {
	var b bool = true

	fmt.Println(reflect.TypeOf(b))

	var s string = strconv.FormatBool(true)
	fmt.Println(reflect.TypeOf(s))
}
