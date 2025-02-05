package main
import (
    "fmt"
    "strings"
)

func main() {
    fmt.Println("strings.Contains(\"test\",\"es\")         = ", strings.Contains("test","es"))
    fmt.Println("strings.Count(\"test\",\"t\")             = ", strings.Count("test","t"))
    fmt.Println("strings.HasPrefix(\"test\",\"te\")        = ", strings.HasPrefix("test","te"))
    fmt.Println("strings.HasSuffix(\"test\",\"st\")        = ", strings.HasSuffix("test","st"))
    fmt.Println("strings.Index(\"test\",\"e\")             = ", strings.Index("test","e"))
    fmt.Println("strings.Join([]string{\"a\",\"b\"},\"-\") = ", strings.Join([]string{"a","b"},"-"))
    fmt.Println("strings.Repeat(\"Q\",3)                   = ", strings.Repeat("Q",3))
    fmt.Println("strings.Replace(\"aaaaa\",\"a\",\"b\",3)  = ", strings.Replace("aaaaa","a","b",3))
    fmt.Println("strings.Split(\"a-b-c-d-e\",\"-\")        = ", strings.Split("a-b-c-d-e","-"))
    fmt.Println("strings.ToLower(\"TEST\")                 = ", strings.ToLower("TEST"))
    fmt.Println("strings.ToUpper(\"test\")                 = ", strings.ToUpper("test"))

    fmt.Println("=======================================================")
    fmt.Println("String to binary...")
    s := "test"
    arr := []byte(s)
    fmt.Println("string = ", s)
    fmt.Println("bytes  = ", arr)

    fmt.Println("=======================================================")
    fmt.Println("Binary to string...")
    byte_arr := []byte{'t','e','s','t'}
    //byte_arr := []byte{97, 98, 99, 100}  // Another way of declaring byte array
    str := string(byte_arr)
    fmt.Println("bytes  = ", byte_arr)
    fmt.Println("string = ", str)
}

