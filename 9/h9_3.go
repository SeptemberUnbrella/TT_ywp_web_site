package main

import "fmt"

func main() {
	s := `After a backslash,certain single character escapes represent
spectal values
n is a line feed pr new line 
t is a tab`
	fmt.Println(s)
}
