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

func totalArea(shapes ...Shape) float64 {
    var area float64
    for _, s := range shapes {
        area += s.area()
    }
    return area
}

func main() {
    c := Circle{0, 0, 10}
    fmt.Println("Circle area     = ", c.area())
    r := Rectangle{0, 0, 10, 10}
    fmt.Println("Rectangle area  = ", r.area())
    fmt.Println("Total area      = ", totalArea(&c, &r))
    fmt.Println("==================================================")
    fmt.Println("And then some more...")
    c1 := Circle{1, 1, 5}
    r1 := Rectangle{1, 2, 3, 4}
    fmt.Println("Circle1 area    = ", c1.area())
    fmt.Println("Rectangle1 area = ", r1.area())
    fmt.Println("Total area      = ", totalArea(&c, &c1, &r, &r1))
}


