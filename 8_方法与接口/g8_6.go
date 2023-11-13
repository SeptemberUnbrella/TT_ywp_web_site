package main

import (
	"fmt"
)

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
	loan   float32
}

type Employee struct {
	Human
	company string
	money   float32
}

func (h Human) SayHi() {
	fmt.Printf("Hi,I an %s you can call me on %s\n", h.name, h.phone)
}

func (h Human) Sing(lyrics string) {
	fmt.Println("La la la la...", lyrics)
}

func (e Employee) SayHi() {
	fmt.Printf("Hi,I am %s,I work at %s.call me on %s\n", e.name, e.company, e.phone)
}

type Men interface {
	SayHi()
	Sing(lyrics string)
}

func main() {
	mike := Student{Human{"Mkie", 25, "222-222-xxx"}, "MIT", 0.00}
	Paul := Student{Human{"Paul", 26, "111-222-xxx"}, "Harvard", 100}
	sam := Student{Human{"sam", 36, "444-222-xxx"}, "Golang I'nc.", 1000}
	Tom := Student{Human{"Tom", 36, "444-222-xxx"}, "Things Ltd.", 5000}

	var i Men
	i = mike
	fmt.Println("This is Mike,a Student:")
	i.SayHi()
	i.Sing("Noveber rain")
	i = Tom
	fmt.Println("This is Tom, an Employee:")
	i.SayHi()
	i.Sing("Born to be wild")
	fmt.Println("Let's use slice of Men and see what happen")
	x := make([]Men, 3)
	x[0], x[1], x[2] = Paul, sam, mike
	for _, value := range x {
		value.SayHi()
	}
}
