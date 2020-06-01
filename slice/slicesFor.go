// Simple program iterating through slice

package main

import(
    "fmt"
)

func main() {

    sports := []string{"Basketball", "Football", "Baseball"}

    for i:= 0; i < len(sports); i++ {
        fmt.Println(sports[i])
    }

    // Single Value Range
    for i := range sports {
        fmt.Println(i)
    }

    // Double Value Range
    for i, name := range sports {
        fmt.Printf("%s is at index %d\n", name, i)
    } 

    // Double Value Range, Ignore Index by using _
    for _, name := range sports {
        fmt.Println(name)
    }

    sports = append(sports, "Golf")
    fmt.Println(sports)
   
}
