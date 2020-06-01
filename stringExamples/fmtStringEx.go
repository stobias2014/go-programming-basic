// Simple example for fmt

package main

import(
    "fmt"
)

func main() {

   number := 1
   
   // Converts number to String
   string := fmt.Sprintf("%d", number)
   
   fmt.Printf("string = %v, is type of %T\n", string, string)
   
   // %q add quotation marks
   fmt.Printf("string = %q, is type of %T\n", string, string)

}
