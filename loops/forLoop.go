// Small program to show for loop in Go

package main 


import(
    "fmt"
)


func main() {

    for x := 0; x < 5; x++ {
        fmt.Println(x)
    }
    
    fmt.Println("----------")    
    for x:= 0; x < 5; x++ {
        fmt.Println(x)
        if x == 3 {
            fmt.Println("I am at three and tired, going to stop....")
            break
        }
    }
}
