package main

import (
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func main() {
	fmt.Println("Hello World")
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":5000", nil)
}
