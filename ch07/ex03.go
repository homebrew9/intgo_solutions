/*
    Add a new perimeter method to the Shape interface to calculate the
    perimeter of a shape. Implement the method for Circle and Rectangle.
*/
package main
import ("fmt"; "math")

type Rectangle struct {
    x1, y1, x2, y2 float64
}
type Circle struct {
    x, y, r float64
}
type Shape interface {
    perimeter() float64
}

func distance(x1, y1, x2, y2 float64) float64 {
    a := x1 - x2
    b := y1 - y2
    return math.Sqrt(a*a + b*b)
}
func (r *Rectangle) area() float64 {
    length := distance(r.x1, r.y1, r.x2, r.y1)
    width := distance(r.x1, r.y1, r.x1, r.y2)
    return length * width
}
func (r *Rectangle) perimeter() float64 {
    length := distance(r.x1, r.y1, r.x2, r.y1)
    width := distance(r.x1, r.y1, r.x1, r.y2)
    return 2 * (length + width)
}

func (c *Circle) area() float64 {
    return math.Pi * c.r * c.r
}
func (c *Circle) perimeter() float64 {
    return 2 * math.Pi * c.r
}

func perimeter(s Shape) float64 {
    fmt.Println("Hello, World from perimeter(s Shape) float64")
    return s.perimeter()
}

func main() {
    r := Rectangle{0, 0, 10, 10}
    fmt.Println("Rectangle area      = ", r.area())
    fmt.Println("Rectangle permieter = ", r.perimeter())
    c := Circle{0, 0, 10}
    fmt.Println("Circle area         = ", c.area())
    fmt.Println("Circle perimeter    = ", c.perimeter())
    fmt.Println("==================================================")
    // I don't think I understood the way to call the interface function.
    // We are not passing the shape as input below, so I don't think the
    // function perimeter(s Shape) is being invoked below.
    fmt.Println(r.perimeter())
}


