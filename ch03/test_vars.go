package main
import "fmt"

func main() {
    var x string = "Hello, World!"
    fmt.Println(x)
    x = "Hola, Mundo!"
    fmt.Println(x)
    x += " => Hallo Welt!"
    x += " => नमस्ते दुनिया!"
    fmt.Println(x)
}

