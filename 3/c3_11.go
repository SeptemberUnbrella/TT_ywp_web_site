package main

import "fmt"

func showMemoryAddress(x string) {
	fmt.Println(&x)
	return
}

func main() {
	var s string = "hello world"
	fmt.Println(&s)
	showMemoryAddress(s)
}
