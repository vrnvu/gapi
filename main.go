package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("debug: running gapi on port 8080")
	s := NewServer()
	log.Fatal(http.ListenAndServe(":8080", s.Router()))
}
