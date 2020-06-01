// Simple program to show how functions work in Go

package main

import(
    "fmt"
)

func add(a int, b int) int {
    return a + b
}

func subtract(a int, b int) int {
    return a - b
}

func multiply(a int, b int) int {
    return a * b
}

func divide(a int, b int) int {
    return a / b
}

func main() {
    a := 45
    b := 5
    
    fmt.Printf("a = %v, b = %v\n", a, b)
    
    fmt.Printf("addition = %v\n", add(a, b))
    fmt.Printf("subtraction = %v\n", subtract(a, b))
    fmt.Printf("multiply = %v\n", multiply(a, b))
    fmt.Printf("division = %v\n", divide(a, b))

}
