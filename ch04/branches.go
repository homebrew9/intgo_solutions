package main
import "fmt"

func main() {
    i := 1
    for i <= 10 {
        if i % 2 == 0 {
            fmt.Println(i, ": Even")
        } else {
            fmt.Println(i, ": Odd")
        }
        i += 1
    }
    fmt.Println("=====================================")
    leap_year_check()
}

func leap_year_check() {
    var input_year int
    var is_leap_year bool = false
    fmt.Print("Enter 4-digit year: ")
    fmt.Scanf("%d", &input_year)
    // A century year is a leap year only if it is divisible by 400
    // Otherwise, the year is leap if it is divisible by 4
    if input_year % 100 == 0 {
        if input_year % 400 == 0 {
            is_leap_year = true
        }
    } else if input_year % 4 == 0 {
        is_leap_year = true
    }
    if is_leap_year {
        fmt.Println(input_year, " is a leap year (it has 366 days)")
    } else {
        fmt.Println(input_year, " is NOT a leap year (it has 365 days)")
    }
}

