// Simple program to show struct in Go

package main

import(
    "fmt"
)

// Definition of Trade struct
type Trade struct {
    Symbol string
    Volume int
    Price float64
    Buy bool
}

func main() {
    // Declaring and assignment trade struct
    t1 := Trade{"Name", 1, 3.45, false}

    fmt.Println(t1)
    
    fmt.Printf("%+v\n", t1)
    fmt.Printf("Name: %v\n", t1.Symbol)    
    fmt.Printf("Volume: %d\n", t1.Volume)    
    fmt.Printf("Price: %.2f\n", t1.Price)    
    fmt.Printf("Buyable: %v\n", t1.Buy)    

    t2 := Trade{
        Symbol: "AMZN",
        Volume: 2,
        Price: 4.56,
        Buy: true,
    }

    fmt.Printf("%+v\n", t2)
}
