package main

import "fmt"

func isEven(i int) bool {
	return i%2 == 0
}

func main() {

	fmt.Printf("%v\n", isEven(1)) //通过%v输出原型
	fmt.Printf("%v\n", isEven(2))
}
