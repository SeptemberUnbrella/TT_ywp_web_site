package main

import (
	"fmt"
)

func sayHello(s string) string { //接收一个string的参数，斌且返回值也是string类型
	return "Hello " + s
}

func main() {
	fmt.Println(sayHello("Gerogo"))
}
