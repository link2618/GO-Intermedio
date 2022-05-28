package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

// De esta manera aplicamos la composicion sobre la herencia
type FullTimeEmployee struct {
	Person
	Employee
	endDate string
}

type TemporatyEmployee struct {
	Person
	Employee
	taxRate int
}

// Creamos la interfaz
type PrintInfo interface {
	getMessage() string
}

// Implementamos la interfaz
func (ftEmployee FullTimeEmployee) getMessage() string {
	return "Full Time Employee"
}

func (tEmployee TemporatyEmployee) getMessage() string {
	return "Temporary Employee"
}

// Creamos el metodo de la interfaz
func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Pepito"
	ftEmployee.age = 27
	ftEmployee.id = 1
	fmt.Printf("%v\n", ftEmployee)
	tEmployee := TemporatyEmployee{}
	getMessage(tEmployee)
	getMessage(ftEmployee)
}
