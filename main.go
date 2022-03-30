package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello, world")
		d, _ := ioutil.ReadAll(r.Body)
		fmt.Fprintf(w, "Hello %s", d)
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Goodbye, world")
	})

	http.ListenAndServe(":9090", nil)
}
