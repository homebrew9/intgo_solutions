/*
    Wrong way to program. The function totalArea has a major issue - whenever
    we define a new shape, we have to change our function to handle it (a
    third parameter for Polygons, a fourth for Squares, etc.)
    It does not use the interface Shape, which is designed to solve problems
    like this. Since both of our shapes have an area method, they both implement
    the Shape interface and we can use cleaner code for totalArea function.
    See the program "interfaces_2.go" for the better program.
*/
package main
import ("fmt"; "math")

type Shape interface {
    area() float64
}
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

func totalArea(circles []Circle, rectangles []Rectangle) float64 {
    var total float64
    for _, c := range circles {
        total += c.area()
    }
    for _, r := range rectangles {
        total += r.area()
    }
    return total
}

func main() {
    c := Circle{0, 0, 10}
    fmt.Println("Circle area    = ", c.area())
    r := Rectangle{0, 0, 10, 10}
    fmt.Println("Rectangle area = ", r.area())
    cir := []Circle{c}
    rect := []Rectangle{r}
    fmt.Println("Total area     = ", totalArea(cir, rect))
}

