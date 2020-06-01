// Simple program to show strings in Go language

package main

import(
    "fmt"
)

func main() {

    bookTitle := "Catcher in the Rye"
    
    fmt.Printf("bookTitle %v\n", bookTitle)
    
    // length in bytes
    fmt.Printf("length of bookTitle: %v\n", len(bookTitle))
    
    fmt.Printf("bookTitle[0]: %v, is type of %T\n", bookTitle[0], bookTitle[0])
    
    // strings in go are immutable
    // cannot assign new value

    // Slice (start, end), 0 based, 1/2 empy range
    fmt.Println(bookTitle[4:8])

    // Slice (no end)
    fmt.Println(bookTitle[4:])

    // Slice (no start)
    fmt.Println(bookTitle[:6])
    
    // Use of + to concatanate
    fmt.Println("book title: " + bookTitle)
}
