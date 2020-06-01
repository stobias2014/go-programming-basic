// Simple program for fizz buzz


package main

import(
    "fmt"
)


func main() {

// print numbers from 1 to 20
// if number is divisible by 3, print fizz
// if number is divisible by 5, print buzz
// if number is divisible by both 3 and 5, print fizz buzz

for x := 0; x < 21; x++ {
    
    if x % 3 == 0 && x % 5 == 0 {
        fmt.Println("fizz buzz")
    } else if x % 5 == 0 {
        fmt.Println("buzz")
    } else if x % 3 == 0 {
        fmt.Println("fizz")   
    } else {
        fmt.Println(x)   
    }

}






}
