package main

import (
	"fmt"
)

type Employee struct {
	Name   string
	Age    int
	Salary float64
}

func main() {
	fmt.Println("vim-go")

	name := "I am a string"
	number := 5
	decimalNumber := 5.3
	isDrinkable := false
	numbers := []int{1, 2, 3, 4, 5}
	employees := map[int]string{
		1: "saul",
		2: "kareli",
		3: "brenda",
	}
	// store address value of name
	address := &name

	employee := Employee{
		Name:   "saul",
		Age:    3,
		Salary: 4500.00,
	}

	fmt.Printf("Name: %s is a type of %T\n", name, name)
	fmt.Printf("Number: %d is a type of %T\n", number, number)
	fmt.Printf("Number with Decimal: %f is a type of %T\n", decimalNumber, decimalNumber)
	fmt.Printf("Can you drink? %t is a type of %T\n", isDrinkable, isDrinkable)
	fmt.Printf("Numbers: %v is a type of %T\n", numbers, numbers)
	fmt.Printf("Employees: %v is a type of %T\n", employees, employees)
	fmt.Printf("address to variable name: %p is a type of %T\n", address, address)
	fmt.Printf("actual value of variable name using pointer: %v\n", *address)
	fmt.Printf("Employee: %+v is type of %T\n", employee, employee)
}
