package main

import (
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	//switch r.Header.Get("Accept") {
	//case "application/json":
	//	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//	w.Write([]byte(`{"message":"Hello world"}`))
	//case "application/xml":
	//	w.Header().Set("Content-Type", "application/xml; charset=utf-8")
	//	w.Write([]byte(`<?xml version="1.0" encoding="utf-8"?><Message>Hello world</Message>`))
	//
	//}
	//w.Header().Set("X-MY-Header", "I am setting a header!")
	//w.Header().Set("X-Forwarded-Host", "host:tttwp")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write([]byte(`{"name":"李路通","age":"20"}`))
}

//func users(w http.ResponseWriter, r *http.Request) {
//	w.Write([]byte("this is users"))
//}

func main() {

	http.HandleFunc("/", helloWorld)
	//	http.HandleFunc("/users/", users)
	http.ListenAndServe("10.0.0.1:8000", nil)

}
