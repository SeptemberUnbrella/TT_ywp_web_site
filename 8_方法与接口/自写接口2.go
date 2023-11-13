package main

import (
	"fmt"
)

type Ainterface interface {
	Say()
}
type Stu struct {
	Name string
}

func (stu Stu) Say() {
	fmt.Println("stu.say()")
}

func main() {

	var stu Stu
	var a Ainterface = stu
	a.Say()
}
