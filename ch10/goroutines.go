package main
import (
    "fmt"
    "strings"
    "time"
    "math/rand"
)

func f(n int) {
    for i := 0; i < 10; i++ {
        fmt.Println(strings.Repeat(">",n), n, ":", i)
    }
}

func f_sim(n int) {
    for i := 0; i < 10; i++ {
        fmt.Println(strings.Repeat(">",n), n, ":", i)
        amt := time.Duration(rand.Intn(250))
        time.Sleep(time.Millisecond * amt)
    }
}

func main() {
    go f(0)
    // The Scanln function has been included because without it, the program
    // would exit before being given the opportunity to print all the numbers.
    var input string
    fmt.Scanln(&input)

    // Goroutines are lightweight and we can easily create thousands of them.
    fmt.Println(strings.Repeat("=",80))
    fmt.Println("Now running ten goroutines...")
    fmt.Println(strings.Repeat("=",80))
    for i := 0; i < 10; i++ {
        go f(i)
    }
    fmt.Scanln(&input)

    fmt.Println(strings.Repeat("=",80))
    fmt.Println("Tweak to make the goroutines run simultaneously...")
    fmt.Println(strings.Repeat("=",80))
    for i := 0; i < 10; i++ {
        go f_sim(i)
    }
    fmt.Scanln(&input)
}

