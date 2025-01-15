package main
import (
    "fmt"
    "time"
)

func Sleep(duration time.Duration) {
    <-time.After(duration)
}

func main() {
    // Time duration is in nanoseconds
    // The value below is for 5 seconds
    var d time.Duration = 5 * 1000000000

    fmt.Println("Before sleep...")
    Sleep(d)
    fmt.Println("After  sleep...")
}

