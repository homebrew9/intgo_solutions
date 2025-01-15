package main
import "fmt"

func makeOddGenerator() func() uint {
    i := uint(1)
    return func() (ret uint) {
        ret = i
        i += 2
        return
    }
}

func main() {
    x := 0
    increment := func() int {
        x++
        return x
    }
    fmt.Println(increment())
    fmt.Println(increment())
    // A function like this together with the nonlocal variables it references
    // is known as a closure. In this case, increment and the variable x form
    // a closure.
    fmt.Println("==================================================")
    fmt.Println("Testing function makeOddGenerator()...")
    nextOdd := makeOddGenerator()
    fmt.Println(nextOdd()) // 1
    fmt.Println(nextOdd()) // 3
    fmt.Println(nextOdd()) // 5
}

