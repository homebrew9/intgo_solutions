package main
import (
    "fmt"
    "time"
)

func fmt_duration(d time.Duration) string {
    d = d.Round(time.Second)
    h := d / time.Hour
    d -= h * time.Hour
    m := d / time.Minute
    d -= m * time.Minute
    s := d / time.Second
    return fmt.Sprintf("%02d:%02d:%02d",h,m,s)
}

func fibo(n uint) uint {
    if n <= 1 {
        return 1
    }
    return fibo(n-1) + fibo(n-2)
}

func spinner(delay time.Duration) {
    // Copied from "The Go Programming Language".
    // Chapter 8 - Goroutines and Channels
    for {
        for _, r := range `\|/` {
            fmt.Printf("\r%c", r)
            time.Sleep(delay)
        }
    }
}

func main() {
    slice := make([]uint, 51)
    for i := 0; i < len(slice); i++ {
        slice[i] = uint(i)
    }
    for _, value := range slice {
        go spinner(100 * time.Millisecond)
        start := time.Now()
        fibonacci := fibo(value)
        elapsed := time.Since(start)
        fmt.Printf("\rfibo(%2d) = %15d ; Calculation time = %15s ; Fmt calc time = %s\n",value,fibonacci,elapsed,fmt_duration(elapsed))
    }
}

