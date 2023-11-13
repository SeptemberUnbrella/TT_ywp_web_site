package main

import "fmt"

func sumNumbers(numbers ...int) int {
	total := 0
	for _, numbers := range numbers {
		total += numbers
	}
	return total
}

func main() {
	result := sumNumbers(1, 2, 3, 4)
	fmt.Printf("the result is %v\n", result)
}
