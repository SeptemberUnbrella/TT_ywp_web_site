package main

import (
	"fmt"
	"reflect"
)

type Drink struct {
	Name string
	Ice  bool
}

func main() {

	a := Drink{
		Name: "lemonade",
		Ice:  true,
	}
	b := Drink{
		Name: "lemonade",
		Ice:  true,
	}

	if a == b {
		fmt.Println("a and b are the same!")
	}
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))

}
