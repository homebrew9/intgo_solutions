package main
import "fmt"

func main() {
    var fahrenheit, celsius float64
    fmt.Print("Enter Fahrenheit value: ")
    fmt.Scanf("%f", &fahrenheit)
    celsius = (fahrenheit - 32) * 5/9
    fmt.Println("Fahrenheit value = ", fahrenheit)
    fmt.Println("Celsius value    = ", celsius)
}

