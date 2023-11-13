package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	URL := "https://ifconfig.io/aa"
	response, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	boby, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", boby)
}
