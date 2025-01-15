package main
import "fmt"

func half(x int) (int, bool) {
    half_value := int(x/2)
    is_even := false
    if x % 2 == 0 {
        is_even = true
    }
    return half_value, is_even
}

func main() {
    for i := 1; i <= 10; i++ {
        x, y := half(i)
        fmt.Println("half(", i, ") => ", x, y)
    }
}

