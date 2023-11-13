package main

import "fmt"

func main() {

	s := "After a backslash,certain single character escapes represent" +
		"sepecial values\nn is a line feed or new line \n\t t is a tab"
	fmt.Println(s)
}
