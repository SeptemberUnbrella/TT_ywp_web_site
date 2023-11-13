package main

import "fmt"

type Sex struct {
	sex string
}

func NewSex(i int) Sex {
	var e Sex
	if i == 0 {
		e.sex = "女"
	} else {
		e.sex = "男"
	}
	return e
}

func main() {
	var i int = 1
	fmt.Printf("%+v\n", NewSex(i))

}
