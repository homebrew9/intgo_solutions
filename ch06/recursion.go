package main
import "fmt"

func factorial(x uint) uint {
    if x == 0 {
        return 1
    }
    return x * factorial(x-1)
}

func main() {
    slice := make([]uint, 20)
    for i := 1; i <= len(slice); i++ {
        slice[i-1] = uint(i)
    }
    fmt.Println(slice)
    for _, val := range slice {
        fmt.Printf("%2d! = %25d\n",val,factorial(val))
    }
}

