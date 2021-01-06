package main

// Is it a good idea to use mux?
// Future API requirements auth and cors
// We could use psql, mdb or json files
// I like json files for starters
import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Server interface exposes a http.Handler
type Server interface {
	Router() http.Handler
}

// api is our router. The idea is to expose the interface and keep this struct private
type api struct {
	router http.Handler
}

func (a *api) Router() http.Handler {
	return a.router
}

// NewServer creates an http.Handler server
func NewServer() Server {
	a := &api{}
	r := mux.NewRouter()

	r.HandleFunc("/ini", a.ini).Methods(http.MethodGet)
	r.HandleFunc("/del", a.del).Methods(http.MethodGet)

	r.HandleFunc("/employees", a.fetchEmployees).Methods(http.MethodGet)

	// We capture name and define a valid regex condition
	r.HandleFunc("/employee/{name:[a-z]+}", a.fetchEmployee).Methods(http.MethodGet)

	r.HandleFunc("/employee", a.createEmployee).Methods(http.MethodPost)

	a.router = r
	return a
}

// TODO mutex or some sort of sync
func (a *api) ini(w http.ResponseWriter, r *http.Request) {
	employees := []Employee{
		{"a", 100, 10},
		{"b", 200, 20},
		{"c", 300, 30},
		{"d", 400, 40},
		{"e", 500, 50},
	}
	WriteEmployees(employees)
}

// TODO mutex or some sort of sync
func (a *api) del(w http.ResponseWriter, r *http.Request) {
	WriteEmployees(make([]Employee, 0))
}

// TODO mutex or some sort of sync
func (a *api) fetchEmployees(w http.ResponseWriter, r *http.Request) {
	employees := ReadEmployees()
	fmt.Println("debug:", employees)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

// TODO mutex or some sort of sync
func (a *api) fetchEmployee(w http.ResponseWriter, r *http.Request) {
	// Returns values captured in the request URL
	// vars is a dictionary whose key-value pairs are variables
	vars := mux.Vars(r)
	fmt.Println("debug:", vars)

	employees := ReadEmployees()

	employee, err := FindEmployee(vars["name"], employees)
	if err != nil {
		fmt.Println("debug:", err)
	}
	fmt.Println("debug:", employee)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employee)
}

// TODO mutex or some sort of sync
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

	employees := ReadEmployees()
	employees = append(employees, employee)
	WriteEmployees(employees)

	w.WriteHeader(http.StatusCreated)
}
