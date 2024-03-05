package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello-world", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK!")
	})
	http.ListenAndServe(":8890", nil)
}

type Person struct {
	Name   string
	Sex    string
	IdType string
	IdNo   string
}
