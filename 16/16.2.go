package main

import (
	"errors"
	"log"
)

func main() {
	var errFatal = errors.New("we onl just error")
	log.Fatal(errFatal)
}
