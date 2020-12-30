package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
)

// Fields must start with capital letters to be exported
type Employee struct {
	Name string `json:"name"`
	Salary int `json:"salary"`
	Sales int `json:"sales"`
}

// NOTE I don't like this api design
type api struct {
	router http.Handler
}

type Server interface {
	Router() http.Handler
}

// Create a new server with two endpoints for our employees
func NewServer() Server {
	a := &api{}
	r := mux.NewRouter()

	r.HandleFunc("/employees", a.fetchEmployees).Methods(http.MethodGet)
	r.HandleFunc("/employee/{name:[a-z]+}", a.fetchEmployee).Methods(http.MethodGet)

	a.router = r
	return a
}

// NOTE unnecessary
func (a *api) Router() http.Handler {
	return a.router
}

func (a *api) fetchEmployees(w http.ResponseWriter, r *http.Request) {
	employees := []Employee{
		Employee{"a", 100, 10},
		Employee{"b", 200, 20},
		Employee{"c", 300, 30},
	}
	fmt.Println("debug:", employees)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func (a *api) fetchEmployee(w http.ResponseWriter, r *http.Request) {

	// Returns values captured in the request URL
	// vars is a dictionary whose key-value pairs are variables
	vars := mux.Vars(r)
	fmt.Println("debug:", vars)

	employee := Employee{"a", 100, 10}
	fmt.Println("debug:", employee)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employee)
}
