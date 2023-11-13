package main

import "fmt"

func main() {
	var cheeses = make([]string, 3)
	cheeses[0] = "Mariolles"
	cheeses[1] = "Epoisses de Bourogone"
	cheeses[2] = "camebert"
	fmt.Println(len(cheeses))
	fmt.Println(cheeses)
	cheeses = append(cheeses[:2], cheeses[2+1:]...) //就是通过覆盖
	fmt.Println(len(cheeses))
	fmt.Println(cheeses)
}
