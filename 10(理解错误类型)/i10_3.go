package main

import (
	"fmt"
)

func main() {
	name, role := "Ricahrd Jupp", "Drummer"
	err := fmt.Errorf("the %v %v quit", role, name)
	if err != nil {
		fmt.Println(err)

	}
}
