package main
import ("fmt"; "math")

type Circle struct {
    x, y, r float64
}

func circleArea(c *Circle) float64 {
    return math.Pi * c.r * c.r
}

func main() {
    c := Circle{0, 0, 10}
    fmt.Printf("Circle center   = (%f, %f)\n", c.x, c.y)
    fmt.Printf("Circle radius   = %f\n", c.r)
    fmt.Println("Area of circle = ", circleArea(&c))
}

