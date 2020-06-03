// Simple program showing defer to make sure resource is closed
// Go has garbage collector

package main

import(
    "fmt"
)

func cleanup(name string) {
    fmt.Printf("Cleaning up: %s\n", name)
}

// Use to clean up and close resource
// defer is used to cleanup
// inside worker is actual code to run, defer is called after
func worker() {
    defer cleanup("Saul Tobias")

    fmt.Println("I am executed before defer")   
}

func main() {
    worker()
}
