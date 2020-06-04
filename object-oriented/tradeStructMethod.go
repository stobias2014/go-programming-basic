// Simple program to show an object with a method


package main

import(
    "fmt"
)

// Definition of a Trade struct
type Trade struct{
    Symbol string
    Volume int
    Price float64
    Buy bool
}

// Defining a method for the Trade struct
// a fucntion with a reference to Trade
// With name of Value
// return a type float64

// (t *Trade) is a receiver
// pointer to Trade struct
// use pointer receiver so we dont get a copy of value
func (t *Trade) Value() float64{
    value := float64(t.Volume) * t.Price
    if t.Buy {
        value = -value
    }
    
    return value
}

func main() {
    // Declaring and initializing a Trade struct
    t1 := Trade{
        Symbol: "AMZN",
        Volume: 3,
        Price: 3.45,
        Buy: true,
    }

   // Using the method of Trade
   fmt.Printf("Value: %.2f\n", t1.Value()) 

}
