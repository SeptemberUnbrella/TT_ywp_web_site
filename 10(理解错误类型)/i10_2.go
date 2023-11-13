package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("something went wrong") //创建一个错误消息
	if err != nil {
		fmt.Println(err)
	}
}
