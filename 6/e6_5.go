package main

import "fmt"

func main() {
	var chesses = make([]string, 3)
	chesses[0] = "Mariollles"
	chesses[1] = "Epoisses de bourgogone"
	var smellyCheeses = make([]string, 2)
	copy(smellyCheeses, chesses[:1])
	smellyCheeses = append(smellyCheeses, chesses[0])
	fmt.Println(smellyCheeses)
}
