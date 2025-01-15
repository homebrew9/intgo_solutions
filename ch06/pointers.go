package main
import "fmt"

func zero_pbv(x int) {
    x = 0
    fmt.Println("Inside zero_pbv, x = ",x)
}

func zero_pbr(xPtr *int) {
    *xPtr = 0
    fmt.Println("Inside zero_pbr:")
    fmt.Println("  => Address of x = ",&xPtr)
    fmt.Println("  => Value of x   = ",*xPtr)
}

func one(xPtr *int) {
    *xPtr = 1
}

func main() {
    fmt.Println("Pass by value...")
    x := 5
    fmt.Println("Before zero_pbv, x = ",x)
    zero_pbv(x)
    fmt.Println("After zero_pbv, x  = ",x)
    fmt.Println("==================================================")
    fmt.Println("Pass by reference...")
    fmt.Println("Before zero_pbr, x = ",x)
    zero_pbr(&x)
    fmt.Println("After zero_pbr, x  = ",x)
    fmt.Println("==================================================")
    fmt.Println("Using new function...")
    // Very important note: new() takes a type as argument, allocates enough
    // memory to fit a value of that type, and return a **pointer** to it.
    xPtr := new(int)
    fmt.Println("Before one, x = ",*xPtr)
    one(xPtr)
    fmt.Println("After one, x  = ",*xPtr)
}

