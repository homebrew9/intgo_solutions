package main
import "fmt"

func main() {
    i := 1
    for i <= 10 {
        fmt.Println(i)
        i += 1
    }
    fmt.Println("==========================")
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }
}

