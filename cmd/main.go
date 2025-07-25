package main

import (
	"fmt"
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, i love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
