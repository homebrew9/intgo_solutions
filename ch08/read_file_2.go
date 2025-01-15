package main
import (
    "fmt"
    "io/ioutil"
)

func main() {
    bs, err := ioutil.ReadFile("poem_rfrost.txt")
    if err != nil {
        return
    }
    str := string(bs)
    fmt.Printf("[%s]", str)
}

