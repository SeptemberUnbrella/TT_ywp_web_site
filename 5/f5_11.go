package main

import "fmt"

func main() {
	defer fmt.Println("I am the first defer")
	defer fmt.Println("I am the second defer")
	defer fmt.Println("I am the third defer")
	fmt.Println("Hello world!")
}

//当出现多条defer语句时，执行顺序是由下而上的
