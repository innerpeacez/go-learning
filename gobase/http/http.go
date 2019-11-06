package main

import (
	"fmt"
	"log"
	"net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

type String string
type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func main() {
	var h Hello
	err := http.ListenAndServe("localhost:4000", h)

	if err != nil {
		log.Fatal(err)
	}
}
