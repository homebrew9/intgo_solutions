package main
import "fmt"

func main() {
    fmt.Print("Enter a number: ")
    var input float64
    fmt.Scanf("%f", &input)
    output := input * 2
    fmt.Println("Number entered   = ", input)
    fmt.Println("Double the value = ", output)
}

