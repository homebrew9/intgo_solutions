package main
import "fmt"

func fib(n int) int {
    if n <= 1 {
        return 1
    }
    return fib(n-1) + fib(n-2)
}

func main() {
    //slice := []int{0,1,2,3,19}
    slice := make([]int, 31)
    for i := 0; i < len(slice); i++ {
        slice[i] = i
    }
    for _,value := range slice {
        fmt.Printf("i = %2d ; fib(%2d) = %10d\n", value, value, fib(value))
    }
}

