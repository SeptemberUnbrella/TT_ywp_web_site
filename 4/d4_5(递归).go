package main

import "fmt"

func feedme(portion int, eaten int) int {
	eaten = portion + eaten
	if eaten >= 5 {
		fmt.Printf("I'm full!I've eaten %d\n", eaten)
		return eaten
	}
	fmt.Printf("I'm still hungry!I've eaten %d\n", eaten)
	return feedme(portion, eaten)
}
func main() {
	fmt.Println(feedme(1, 0))
}
