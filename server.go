package main

// Is it a good idea to use mux?
// Future API requirements auth and cors
// We could use psql, mdb or json files
// I like json files for starters
import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
)

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

	// We capture name and define a valid regex condition
	r.HandleFunc("/employee/{name:[a-z]+}", a.fetchEmployee).Methods(http.MethodGet)


	r.HandleFunc("/employee", a.createEmployee).Methods(http.MethodPost)

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


func (a *api) createEmployee(w http.ResponseWriter, r *http.Request) {
	// curl --header "Content-Type: application/json" --request POST --data '{"name":"xyz","salary":1500, "sales":30}' http://localhost:8080/employee
	// We attempt to unmarshall our r.Body into an Employee
	var employee Employee
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error unmarshalling request body")
		return
	}
	fmt.Println("debug:", employee)
	w.WriteHeader(http.StatusCreated)
}
