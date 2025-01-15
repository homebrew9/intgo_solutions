/*
    Here, "area()" is an example of a method.
    In between the keyword func and the function name, we add a "receiver".
    The receiver is like a parameter - it has a name and a type - but by
    creating the function in this way, it allows us to call the function
    using the . operator => r.area()
*/
package main
import ("fmt"; "math")

type Rectangle struct {
    x1, y1, x2, y2 float64
}
type Circle struct {
    x, y, r float64
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
func (c *Circle) area() float64 {
    return math.Pi * c.r * c.r
}

func main() {
    r := Rectangle{0, 0, 10, 10}
    fmt.Println("Rectangle area = ", r.area())
    c := Circle{0, 0, 10}
    fmt.Println("Circle area    = ", c.area())
}

