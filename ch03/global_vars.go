package main
import "fmt"

var x string = "Hello, World!"

func main() {
    // Local variables and constants take precendence
    const x string = "I can't change!"
    fmt.Println(x)
    fmt.Println("In main() function, x =", x)
    f()
    g()
}

func f() {
    fmt.Println("In f() function, x =", x)
}

func g() {
    fmt.Println("In g() function, x =", x)
}

