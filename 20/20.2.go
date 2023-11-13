package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	Name    string   `json:"name,omitempty"`
	Age     int      `json:"age,omitempty"`
	Hobbies []string `json:"hobbies,omitempty"`
}

func main() {
	jsonStingData := `{"name":"George","age":40,"hobbies":["Cycling","abc"]}`
	jsonByteData := []byte(jsonStingData)
	p := person{}
	err := json.Unmarshal(jsonByteData, &p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", p)
}
