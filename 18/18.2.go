package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func get_mothed(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case "GET":
		for k, v := range r.URL.Query() {
			fmt.Printf("%s: %s\n", k, v)
		}
		w.Write([]byte("Received GET request!"))
	case "POST":
		reqBoby, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", reqBoby)
		w.Write([]byte("Received POST request!"))
	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotFound) + "\n"))
	}
}
func main() {
	http.HandleFunc("/", get_mothed)
	http.ListenAndServe("10.0.0.1:8080", nil)

}
