// Simple program to show how errors work in Go

package main

import(
    "fmt";
    "math"
)

// Go functions can return more than one value
// nil is no value
func sqrt(number float64) (float64, error) {
    if number < 0 {
        return 0.0, fmt.Errorf("Cannot have a square root of negative number [%f]", number)   
    } else {
        return math.Sqrt(number), nil   
    }
}


func main() {

    // calling the function
    result, error := sqrt(9.0)
    
    // if error exists, print error message

    if error != nil {
        fmt.Printf("ERROR: %s\n", error)
    } else {
        fmt.Printf("square root = %f\n", result)
    }

    fmt.Println("--------------")

    // calling the function with value that causes error
    r, error := sqrt(-9.0)
    
    if error != nil {
        fmt.Printf("ERROR: %s\n", error)
    } else {
        fmt.Printf("square root = %f\n", r)
    }
}
