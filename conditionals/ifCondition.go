// Simple program to demonstrate the use of if condition

package main

import (
    "fmt"
)

func main() {

    x := 25

    if x > 20 {
        fmt.Printf("%v is greater than 20\n", x)
    } 

    // if-else condition
    if x < 100 {
        fmt.Printf("%v is less than 100\n", x)
    } else {
        fmt.Printf("%v is greater than a 100\n", x)
    }

    // logical operator for AND
    if x > 20 && x < 30 {
        fmt.Printf("%v is between 20 and 30\n", x)
    }
    
    // logical operator for OR
    if x < 20 || x > 30 {
        fmt.Printf("%v is out of range\n", x)
    }

    // define some value from within condition
    a := 11.0
    b := 20.0

    if frac := a / b; frac > 0.5 {
        fmt.Printf("%v is greater than 0.5\n", frac)
    }

   
}
