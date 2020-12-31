package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)


// Fields must start with capital letters to be exported
// Name is id
type Employee struct {
	Name string `json:"name"`
	Salary int `json:"salary"`
	Sales int `json:"sales"`
}

const file = "employees.json"

func ReadEmployees() []Employee {
	jsonFile, err := os.Open(file)
	defer jsonFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var employees []Employee

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &employees)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	return employees
}

func WriteEmployees(employees []Employee) {
	toWrite, _ := json.Marshal(employees)

	// TODO What mask is 0644 exactly?
	err := ioutil.WriteFile(file, toWrite, 0644)
	if err != nil {
		panic(err)
	}
}

// Return true if both slices contain the same employees
func Equals(xs, ys []Employee) bool {
	if len(xs) != len(ys) {
		return false
	}

	m := make(map[string]bool)

	for _, x := range xs {
		m[x.Name] = true
	}

	for _, y := range ys {
		if _, exists := m[y.Name]; !exists {
			return false
		}
	}

	return true
}
