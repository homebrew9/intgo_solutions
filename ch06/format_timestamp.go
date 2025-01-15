/* A quick and dirty script to test the unusual datetime formats of Golang.
 * I have used fmt.Scanln() function, which is not the best choice, because:
 * a) if there are spaces in input, then the last token is accepted
 * b) it does not recognize empty string input
 * Some day I will use a more robust input accepting function.
 */
package main
import (
    "fmt"
    "time"
)

func get_line(s string, n int) string {
    var line string = ""
    for i := 1; i <= n; i++ {
        line += s
    }
    return line
}

func format_now(fmt string) (string,string) {
    t := time.Now()
    return t.String(), t.Format(fmt)
}

func main() {
    var ts_format string
    for {
        fmt.Println("Enter date format (q to quit): ")
        fmt.Scanln(&ts_format)
        //fmt.Printf("You entered: [%s]\n", ts_format)
        if ts_format == "q" {
            break
        }
        unfmt_ts, fmt_ts := format_now(ts_format)
        fmt.Printf("Unformatted current timestamp : [%s]\n", unfmt_ts)
        fmt.Printf("Timestamp format              : [%s]\n", ts_format)
        fmt.Printf("Formatted current timestamp   : [%s]\n", fmt_ts)
        fmt.Println(get_line("=",90))
    }
}

