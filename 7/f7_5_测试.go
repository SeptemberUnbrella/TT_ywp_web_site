package main

import "fmt"

type Student struct {
	Name    string
	Age     int
	Subject Subject
}
type Subject struct {
	class_one   string
	class_two   string
	class_three string
}

func main() {
	lilutong := Student{
		Name: "李路通",
		Age:  22,
		Subject: Subject{
			class_one:   "语文",
			class_three: "数学",
			class_two:   "英语",
		},
	}
	fmt.Printf("%+v\n", lilutong)
	fmt.Println(lilutong.Name)
}
