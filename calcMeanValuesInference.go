// Simple program to calculate mean value using type inference

package main

import (
    "fmt"
)

func main() {
    // the use assignment operator with colon tells Go to infer value
    x := 3.0
    y := 2.0

    fmt.Printf("x=%v, with the type infered of %T\n", x, x)
    fmt.Printf("y=%v, with the type infered of %T\n", y, y)

    mean := (x + y) / 2.0
    fmt.Printf("mean=%v, with the type infered of %T\n", mean, mean)
}
