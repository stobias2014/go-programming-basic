// Simple program to show parameter passing

package main

import(
    "fmt"
)

func addTwoAtIndex(values []int, i int) {
    values[i] += 2
}

func doubleAtIndex(values []int, i int) {
   values[i] *= 2
}

func doubleNumber(num int) {
    num += 2
}

// passing a pointer to n
// dereference by using *, get the actual value
func doublePtr(n *int) {
    *n *= 2
}

// objects are pass by reference, literals are pass by value
// you can use pointers to access direct value
func main() {
    i := 2
    numbers := []int{1, 2, 3, 4, 5}
    
    fmt.Println("Before any functions applied")
    fmt.Println(numbers)    
    
    fmt.Println("After adding the add function")
    addTwoAtIndex(numbers, i)
    fmt.Println(numbers)
    
    fmt.Println("After multiplication function")
    doubleAtIndex(numbers, 3)
    fmt.Println(numbers)

    num := 5
    fmt.Printf("Number before function: %v\n", num)
    fmt.Println("Using regular function, value is")
    doubleNumber(num)
    fmt.Println(num)
    
    fmt.Println("Using pointer function")
    doublePtr(&num)
    fmt.Println(num)    
        
}

