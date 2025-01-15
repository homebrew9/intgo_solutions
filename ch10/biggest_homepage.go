/*
   Most important tip: In order to make this program work, you must set the
   proxy server correctly. Run the following two commands in the Terminal first.

   set HTTP_PROXY=http://gate.swissre.com:8080
   set HTTPS_PROXY=http://gate.swissre.com:8080
*/
package main
import (
    "fmt"
    "io/ioutil"
    "net/http"
)

type HomePageSize struct {
    URL string
    Size int
}

func main() {
    urls := []string{
        "https://www.apple.com",
        "https://www.amazon.com",
        "https://www.google.com",
        "https://www.microsoft.com",
        "https://www.yahoo.com",
        "https://www.facebook.com",
        "https://www.newegg.com",
    }

    results := make(chan HomePageSize)
    for _, url := range urls {
        go func(url string) {
            res, err := http.Get(url)
            if err != nil {
                panic(err)
            }
            defer res.Body.Close()

            bs,err := ioutil.ReadAll(res.Body)
            if err != nil {
                panic(err)
            }

            results <- HomePageSize{
                URL: url,
                Size: len(bs),
            }
        }(url)
    }

    var biggest HomePageSize
    for range urls {
        result := <-results
        fmt.Printf("%-30s => %9d\n", result.URL, result.Size)
        if result.Size > biggest.Size {
            biggest = result
        }
    }
    fmt.Println("==================================================")
    fmt.Println("The biggest home page:", biggest.URL)
}

