// Simple program to show how to create constructor for objects in Go

package main

import(
    "fmt";
    "os"
)

// Definition of Trade struct
type Trade struct{
    Symbol string
    Volume int
    Price float64
    Buy bool
}

// Function used as Constructor for Trade object
// Takes in the parameters that need initialization inside of Trade
// Validates parameters
// Created a new struct with pointer to a Trade
func NewTrade(symbol string, volume int, price float64, buy bool) (*Trade, error) {
    if symbol == "" {
        return nil, fmt.Errorf("Cannot have an empty name")   
    }

    if volume < 0 {
        return nil, fmt.Errorf("Cannot have a negative volume, %d", volume)
    }
    
    if price < 0 {
        return nil, fmt.Errorf("Cannot have a negative price, %.2f", price)
    }

    trade := &Trade{
        Symbol: symbol,
        Volume: volume,
        Price: price,
        Buy: buy,
    }

    return trade, nil

}

func (t *Trade) Value() float64{
    value := float64 (t.Volume) * t.Price
    if t.Buy {
        value = -value
    }
    return value
}

func main() {
    trade, err := NewTrade("AMZN", 3, 3.45, true)
    
    if err != nil {
        fmt.Printf("Error %s\n", err)
        os.Exit(1)
    }

    fmt.Println(trade.Value())

}
