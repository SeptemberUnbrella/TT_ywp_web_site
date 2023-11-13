package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer bytes.Buffer

	for i := 0; i < 500; i++ {
		buffer.WriteString("李路通\t") //通过使用缓存区来拼接字符串效率更高
	}
	a := buffer.String()
	fmt.Println(a)
}
