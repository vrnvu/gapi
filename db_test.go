package main

import (
	"testing"
)

// defered at every write operation to clear file
func clearFile() {
	WriteEmployees(make([]Employee, 0))
}

func TestReadEmpty(t *testing.T) {
	// Assert employees.json is empty
	// Then read empty slice
	got := ReadEmployees()
	want := make([]Employee, 0)
	if !Equals(got, want) {
		t.Errorf("got %#v, wanted %#v", got, want)
	}
}

func TestWriteOne(t *testing.T) {
	defer clearFile()
	e := Employee{"e0", 1000, 10}
	employees := []Employee{e}
	WriteEmployees(employees)
	got := ReadEmployees()
	want := []Employee{e}
	if !Equals(got, want) {
		t.Errorf("got %#v, wanted %#v", got, want)
	}
}

func TestAppendOne(t *testing.T) {
	defer clearFile()
	e := Employee{"e0", 1000, 10}
	employees := []Employee{e}
	WriteEmployees(employees)

	// read again and append again
	got := ReadEmployees()
	got = append(got, e)
	WriteEmployees(got)
	got = ReadEmployees()
	want := []Employee{e, e}
	if !Equals(got, want) {
		t.Errorf("got %#v, wanted %#v", got, want)
	}
}

func TestFindEmployee(t *testing.T) {
	defer clearFile()
	e0 := Employee{"e0", 1000, 10}
	e1 := Employee{"e1", 1000, 10}
	e2 := Employee{"e2", 2000, 20}
	nameToFind := "e1"
	employees := []Employee{e0, e1, e2}
	WriteEmployees(employees)
	list := ReadEmployees()
	got, err := FindEmployee(nameToFind, list)
	if err != nil {
		t.Errorf("e1 not found in %#v", list)
	}
	if got.Name != nameToFind {
		t.Errorf("FindEmployee function is broken")
	}
}
