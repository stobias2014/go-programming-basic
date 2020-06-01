// Simple program to show slice

package main

import(
    "fmt"
)

func main() {

    // Slice is a sequence of items
    // Items in slice must be of same type
    sports := []string{"Football", "Basketball", "Baseball"}
    
    fmt.Printf("sports: %v, is a type of %T\n", sports ,sports)

    // Length
    fmt.Println(len(sports))

    // Access item by index
    fmt.Println(sports[0])

    // Slices of the sequence
    fmt.Println(sports[1:])

}
