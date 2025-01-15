package main
import "fmt"

func main() {
    var x [5]int
    x[4] = 100
    fmt.Println(x)
    fmt.Println("==================================================")
    var y [5]float64
    y[0] = 98
    y[1] = 93
    y[2] = 77
    y[3] = 82
    y[4] = 83
    fmt.Println("Array y = ", y)
    fmt.Println("Length  = ", len(y))
    var total float64 = 0
    for i := 0; i < len(y); i++ {
        total += y[i]
    }
    // Good information on why Golang avoided implicit type-casting:
    // https://stackoverflow.com/questions/31353522/why-must-i-convert-an-integer-to-a-float64-to-type-match
    // https://blog.golang.org/constants
    // https://golang.org/doc/faq#conversions
    fmt.Println("Average = ", total/float64(len(y)))
    fmt.Println("==================================================")
    fmt.Println("Find average using for..range")
    var new_total float64 = 0
    for _, value := range y {
        new_total += value
    }
    fmt.Println("Average = ", new_total/float64(len(y)))
    fmt.Println("==================================================")
    // Another way to declare an array
    // Either comma or closing brace (or both) must be added after last element on same line!
    z := [5]float64 {
        98,
        93,
        77,
        82,
        83,
    }
    fmt.Println(z)
    // Yet another way
    w := [5]float64{1,2,3,4,5}
    fmt.Println(w)
}

