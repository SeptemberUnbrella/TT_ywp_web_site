package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	files, err := ioutil.ReadDir("D:\\GOPROject\\src\\tt\\21")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fmt.Println(file.Mode(), file.Name(), file.ModTime(),
			file.Size(), file.IsDir())
	}
}
