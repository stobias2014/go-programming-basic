// Simple program to show how maps work


package main

import(
    "fmt"
)

func main() {

    // Declaration and Initialization of map
    // Keys must be of same type
    // Values must be of same type
    stocks := map[string]float64{
        "AMZN": 1345.65,
        "GOOG": 435.67,
        "MSFT": 96.81,  // Must have trainling comma in multi line 
    }
    
    fmt.Println(stocks)
    
    // Number of items
    fmt.Printf("Length: %v \n", len(stocks))

    // Get a Value by specifying Key
    fmt.Printf("AMZN = %v \n", stocks["AMZN"])

    // Get a Zero value if not found
    fmt.Println(stocks["TSLA"])

    // Use of two values to see if found
    // second value (ok) is boolean value
    value, ok := stocks["TSLA"]
    if !ok {
        fmt.Println("TSLA not found. ")
    } else {
        fmt.Println(value)
    }

    // Setting a value
    stocks["TSLA"] = 1023.45
    
    fmt.Println(stocks)

    // Delete an item
    delete(stocks, "AMZN")
    fmt.Println(stocks)

    // Iteration of a map

    // Single value "for" is on the keys
    for key := range stocks {
        fmt.Println(key)
    }

    // Double value "for" is on key and value
    for key, value := range stocks {
        fmt.Printf("Stock: %s --- Value = %.2f\n", key, value)
    }
}
