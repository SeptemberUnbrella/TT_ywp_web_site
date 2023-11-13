package main

import "fmt"

func main() {

	s := "c"
	switch s {
	case "a":
		fmt.Println("The letter a！")
	case "b":
		fmt.Println("the letter b!")
	default:
		fmt.Println("this is default")

	}
}
