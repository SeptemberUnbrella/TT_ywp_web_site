package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type User struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

func main() {
	v := User{}
	response, err := http.Get("http://10.0.0.1:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&v)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", v)
	fmt.Println(strconv.Atoi(v.Age))
}
