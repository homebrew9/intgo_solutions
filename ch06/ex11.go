package main
import "fmt"

func swap(x, y *int) (*int, *int) {
    temp := new(int)
    *temp = *y
    *y = *x
    *x = *temp
    return x, y
}

func swap_better(x, y *int) (x *int, y *int) {
    // This is the answer in the book. So looks like you can do multiple
    // assignments even for pointer values in Golang! If I remember correctly,
    // this is not possible in C.
    *y, *x = *x, *y
    //return x, y
    return
}

func main() {
    x, y := 113, 355
    fmt.Printf("Before swap : (x, y) = (%d, %d)\n", x, y)
    swap(&x, &y)
    fmt.Printf("After swap  : (x, y) = (%d, %d)\n", x, y)
    fmt.Println("==================================================")
    fmt.Println("One more time; use swap_better() this time...")
    fmt.Printf("Before swap : (x, y) = (%d, %d)\n", x, y)
    swap_better(&x, &y)
    fmt.Printf("After swap  : (x, y) = (%d, %d)\n", x, y)
}

