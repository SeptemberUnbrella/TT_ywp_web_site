package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("example03.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 066)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	log.SetOutput(f)
	for i := 1; i <= 5; i++ {
		log.Printf("LOG iteration %d", i)

	}
}
