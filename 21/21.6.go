package main

import (
	"log"
	"os"
)

func main() {
	err := os.Remove("D:\\GOPROject\\src\\tt\\21\\5.txt")
	if err != nil {
		log.Fatal(err)
	}
}
