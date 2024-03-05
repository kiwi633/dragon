package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/hello-world", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK!")
		b := r.Body
		defer b.Close()
		var person Person
		json.NewDecoder(b).Decode(&person)
		fmt.Println(person)
	})
	http.ListenAndServe(":8890", nil)
}

type Person struct {
	Name   string `json:"name"`
	Sex    string `json:"sex"`
	IdType string `json:"idType"`
	IdNo   string `json:"idNo"`
}

type User struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	CreateTime time.Time `json:"createTime"`
}
