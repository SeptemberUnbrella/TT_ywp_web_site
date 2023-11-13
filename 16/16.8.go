package main

import "fmt"

type Animal struct {
	Name  string
	Color string
}

func main() {
	a := Animal{
		Name:  "cat",
		Color: "Black",
	}
	fmt.Printf("%v \n", a)
	fmt.Printf("%+v\n", a)
}
