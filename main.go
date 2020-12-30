package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("running gapi")
	s := NewServer()
	log.Fatal(http.ListenAndServe(":8080", s.Router()))
}
