package main

import "fmt"

func main() {
	var cheeses = make([]string, 2)
	cheeses[0] = "Mairolles"
	cheeses[1] = "Epoisses de Bourgone"
	cheeses = append(cheeses, "1", "2", "3")
	fmt.Println(cheeses)
	fmt.Println(cheeses[3])
}
