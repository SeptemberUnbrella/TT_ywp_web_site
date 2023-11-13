package main

import "fmt"

func main() {
	var cheeses = make([]string, 2)
	cheeses[0] = "Marioles"
	cheeses[1] = "Epoisses de Bourgoone"
	cheeses = append(cheeses, "李路通")
	fmt.Println(cheeses[2])
}
