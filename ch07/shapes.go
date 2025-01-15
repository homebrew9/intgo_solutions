package main
import ("fmt"; "math")

func distance(x1, y1, x2, y2 float64) float64 {
    xdiff := x1 - x2
    ydiff := y1 - y2
    dist := math.Sqrt(math.Pow(xdiff,2) + math.Pow(ydiff,2))
    return dist
}

func rectangle_area(x1, y1, x2, y2 float64) float64 {
    length := distance(x1, y1, x2, y1)
    width := distance(x1, y1, x1, y2)
    return length * width
}

func circle_area(x1, y1, r float64) float64 {
    return math.Pi * r * r
}

func main() {
    var rx1, ry1 float64 = 0, 0
    var rx2, ry2 float64 = 10, 10
    fmt.Printf("Rectangle diagonal: (%f, %f) and (%f, %f)\n", rx1, ry1, rx2, ry2)
    fmt.Printf("Rectangle area    : %f\n", rectangle_area(rx1, ry1, rx2, ry2))
    var cx1, cy1, cr float64 = 0, 0, 10
    fmt.Printf("Circle center and radius: (%f, %f) and %f\n", cx1, cy1, cr)
    fmt.Printf("Circle area             : %f\n", circle_area(cx1, cy1, cr))
}

