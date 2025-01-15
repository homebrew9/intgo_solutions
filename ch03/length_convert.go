package main
import "fmt"

func main() {
    var feet, meters float64
    fmt.Print("Enter feet value: ")
    fmt.Scanf("%f", &feet)
    meters = feet * 0.3048
    fmt.Println("Feet value   = ", feet)
    fmt.Println("Meters value = ", meters)
}

