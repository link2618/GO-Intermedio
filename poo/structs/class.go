package main

import "fmt"

type Employee struct {
	id   int
	name string
}

func (e *Employee) SetID(id int) {
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) GetId() int {
	return e.id
}

func (e *Employee) GetName() string {
	return e.name
}

func main() {
	e := Employee{}

	fmt.Printf("%v\n", e)

	e.id = 1
	e.name = "John"

	fmt.Printf("%v\n", e)

	e.SetID(5)
	e.SetName("Hugo")

	fmt.Println(e.GetId())
	fmt.Println(e.GetName())
}
