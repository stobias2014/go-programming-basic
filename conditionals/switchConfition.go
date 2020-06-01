// Simple program for switch condition

package main

import(
    "fmt"
)

func main() {

   x := 21


    switch x {
        case 18:
            fmt.Println("You are not old enough to drink")
        case 21:
            fmt.Println("You are just at the age to drink")
        case 22:
            fmt.Println("You have been able to drink")
        default:
            fmt.Println("idk ")

    }

    switch {
        case x > 18:
            fmt.Println("You are older than 18")
        case x == 21:
            fmt.Println("You are 21")
        default:
            fmt.Println("You are younger than 18")

    }

}
