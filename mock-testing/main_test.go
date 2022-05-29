package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetFullTimeEmployeeById(t *testing.T) {
	type args struct {
		id  int
		dni string
	}
	tests := []struct {
		name     string
		args     args
		mockFunc func()
		want     FullTimeEmployee
		wantErr  bool
	}{
		// TODO: Add test cases.
		{name: "Test1",
			args: args{id: 1, dni: "552"},
			mockFunc: func() {
				GetEmployeeById = func(id int) (Employee, error) {
					return Employee{
						Id:       1,
						Position: "CEO",
					}, nil
				}

				GetPersonByDNI = func(dni string) (Person, error) {
					return Person{
						DNI:  "552",
						Name: "Jhon",
						Age:  38,
					}, nil
				}
			},
			want: FullTimeEmployee{
				Person: Person{
					DNI:  "552",
					Name: "Jhon",
					Age:  38,
				},
				Employee: Employee{
					Id:       1,
					Position: "CEO",
				},
			},
			wantErr: false},
		{name: "Test2",
			args: args{id: 2, dni: "96D"},
			mockFunc: func() {
				GetEmployeeById = func(id int) (Employee, error) {
					return Employee{
						Id:       2,
						Position: "Manager",
					}, nil
				}

				GetPersonByDNI = func(dni string) (Person, error) {
					return Person{}, fmt.Errorf("Person not found")
				}
			},
			want: FullTimeEmployee{
				Person: Person{
					DNI:  "96D",
					Name: "Marta",
					Age:  29,
				},
				Employee: Employee{
					Id:       2,
					Position: "Manager",
				},
			},
			wantErr: true},
	}

	originalGetEmployeeById := GetEmployeeById
	originalGetPersonByDNI := GetPersonByDNI

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()

			got, err := GetFullTimeEmployeeById(tt.args.id, tt.args.dni)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFullTimeEmployeeById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("GetFullTimeEmployeeById() = %v, want %v", got, tt.want)
				}
			}
		})
	}

	GetEmployeeById = originalGetEmployeeById
	GetPersonByDNI = originalGetPersonByDNI

}
