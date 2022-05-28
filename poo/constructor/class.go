package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{id: id, name: name, vacation: vacation}
}

func main() {
	e := Employee{}
	fmt.Printf("%v\n", e)

	e2 := Employee{id: 12, name: "Name", vacation: true}
	fmt.Printf("%v\n", e2)

	e3 := new(Employee)
	fmt.Printf("%v\n", *e3)

	e4 := NewEmployee(10, "Hugo", true)
	fmt.Printf("%v\n", *e4)
}
