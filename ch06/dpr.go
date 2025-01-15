package main
import "fmt"

func first() {
    fmt.Println("1st")
}

func second() {
    fmt.Println("2nd")
}

func test_defer() {
    defer second()
    first()
    first()
    first()
}

func main() {
    fmt.Println("Testing defer...")
    test_defer()
    fmt.Println("==================================================")
    fmt.Println("Testing panic and recover...")
/*
    panic("PANIC")
    str := recover() // this will never happen, control doesn't reach here!
    fmt.Println(str)
*/
    // Pair the recover() function with defer
    defer func() {
        str := recover()
        fmt.Println(str)
    }()
    panic("PANIC")
}

