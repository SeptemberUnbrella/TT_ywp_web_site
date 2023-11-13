package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file, err := ioutil.ReadFile("foo.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("%s", file)
}
