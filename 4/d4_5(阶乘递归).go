package main

import "fmt"

func jieCheng(i int) int {
	if i == 1 || i == 0 {
		return 1
	} else {
		return i * jieCheng(i-1)
	}
}

func main() {
	fmt.Println(jieCheng(5))
}
