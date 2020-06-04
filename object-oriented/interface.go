// Simple program to show the use of interface in Go

package main

import(
    "fmt"
    "math"
)

// Definition of Square struct
type Square struct{
    Length float64
}

// Definition of method belonging to Square struct
func (s *Square) Area() float64{
    return s.Length * s.Length
}

// Definition of Circle struct
type Circle struct{
    Radius float64
}

// Definition of method belonging to Circle struct
func (c *Circle) Area() float64{
    return math.Pi * c.Radius * c.Radius
}

// add all the areas of a slice of Shapes
func sumAreas(shapes []Shape) float64{
    total := 0.0
    
    for _, shape := range shapes{
        total +=  shape.Area()
    }

    return total
}

// Shape is an interface
type Shape interface{
    Area() float64
}

func main() {

     s := &Square{3.4}
     fmt.Printf("Square Area: %.2f\n", s.Area())

     c := &Circle{4.5}
     fmt.Printf("Circle Area: %.2f\n", c.Area())

     shapes := []Shape{s, c}
     fmt.Printf("Area for all Shapes: %.2f\n", sumAreas(shapes))

}
