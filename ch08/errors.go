package main

import ("fmt"; "errors")

func main() {
    err := errors.New("My custom error message")
    fmt.Println(err)
}

