package main
import "fmt"

func main() {
    var input_int int
    fmt.Print("Enter an integer in range [0-5]: ")
    fmt.Scanf("%d", &input_int)
    switch input_int {
    case 0: fmt.Println("Zero")
    case 1: fmt.Println("One")
    case 2: fmt.Println("Two")
    case 3: fmt.Println("Three")
    case 4: fmt.Println("Four")
    case 5: fmt.Println("Five")
    default: fmt.Println("Outside range")
    }
}

