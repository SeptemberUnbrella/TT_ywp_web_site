package main

import (
	"io/ioutil"
	"log"
)

func main() {
	Write_txt := "This Is TT Web Server"
	err := ioutil.WriteFile("D:\\GOPROject\\src\\tt\\21\\3.txt", []byte(Write_txt), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
