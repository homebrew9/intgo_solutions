package main
import (
    "fmt"
    "time"
)

func pinger(c chan string) {
    for i := 0; ; i++ {
        c <- "ping"
    }
}

func ponger(c chan string) {
    for i := 0; ; i++ {
        c <- "pong"
    }
}

func printer(c chan string) {
    for {
        msg := <- c
        fmt.Println(msg)
        // Alternative way: fmt.Println(<-c)
        time.Sleep(time.Second * 1)
    }
}

func main() {
    var c chan string = make(chan string)

    go pinger(c)
    go ponger(c)
    go printer(c)

    var input string
    fmt.Scanln(&input)
}

