package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	fileBytes, err := ioutil.ReadFile("D:\\GOPROject\\src\\tt\\21\\1.txt")
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(fileBytes)
	fmt.Println(string(fileBytes))
}
