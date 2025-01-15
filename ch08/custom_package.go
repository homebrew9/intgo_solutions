package main

import "fmt"
import "intgo/ch08/math"

func main() {
    //xs := []float64{}
    xs := []float64{1,2,3,4}
    avg := math.Average(xs)
    min := math.Min(xs)
    max := math.Max(xs)
    fmt.Println("slice = ", xs)
    fmt.Println("avg   = ", avg)
    fmt.Println("min   = ", min)
    fmt.Println("max   = ", max)
}

