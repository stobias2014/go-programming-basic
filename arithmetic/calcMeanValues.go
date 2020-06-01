// Simple Go program to calculate the mean of two numbers

package main

import (
    "fmt"
)

func main() {
    
    // Declare variables x and y as integer types
    var x float64
    var y float64
    
    // assign integer values to variables x and y
    x = 1 
    y = 2
    
    // Printing the values of x and y and their type
    fmt.Printf("x=%v, with the type of %T\n", x, x)
    fmt.Printf("y=%v, with the type of %T\n", y, y)

    // Declare variable mean of type integer
    var mean float64
    mean = (x + y) / 2.0
    
    // Printing the value of mean and its type
    fmt.Printf("mean=%v, with the type of %T\n", mean, mean)
}
