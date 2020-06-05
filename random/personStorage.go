package main

import (
	"bufio"
	"fmt"
	"os"
)

type Person struct {
	Name string
	Age  int
}

type PersonStorage struct {
	Persons []Person
}

func NewPerson(name string, age int) (*Person, error) {
	// Check if name is empty
	if name == "" {
		return nil, fmt.Errorf("error: cannot have empty name")
	}

	// Check if age is less than 0
	if age < 0 {
		return nil, fmt.Errorf("error: cannot have age less than 0")
	}

	// Create a reference to Person struct
	person := &Person{
		Name: name,
		Age:  age,
	}

	return person, nil
}

func main() {

	people := &PersonStorage{}

	// Create a reader os Stdin to read
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < 5; i++ {
		fmt.Println("Enter name: ")

		// Read text until delimiter
		name, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Printf("You entered: %s\n", name)

		fmt.Println("Enter age:")

		var age int
		_, err = fmt.Scanf("%d", &age)

		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Printf("You entered: %d\n", age)

		person, err := NewPerson(name, age)
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Printf("\nPerson created: %+v\n", person)

		fmt.Println("Placing in Person Storage")

		people.Persons = append(people.Persons, *person)

		fmt.Printf("\nPerson Storage Contains: %+v\n\n", people.Persons)
	}
}
