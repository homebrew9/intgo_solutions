package main
import (
    "fmt"
    "hash/crc32"
    "strings"
)

func main() {
    // create a hasher
    h := crc32.NewIEEE()
    // write our data to it
    h.Write([]byte("test"))
    // calculate the crc32 checksum
    v := h.Sum32()
    fmt.Println("test")
    fmt.Println(v)
    fmt.Println(strings.Repeat("=",75))
    fmt.Println("CRC32 hashes on all elements of an array...")
    fmt.Println(strings.Repeat("=",75))
    slice := []string{
        "The quick brown fox jumps over the lazy dog",
        "the quick brown fox jumps over the lazy dog",
        "The woods are lovely, dark and deep",
        "Now is the time for all young men",
        "abracadabra",
    }
    for _, value := range(slice) {
        fmt.Println("String       = ", value)
        h.Write([]byte(value))
        v := h.Sum32()
        fmt.Println("CRC32 Hash   = ", v)
        fmt.Println(strings.Repeat("-",65))
    }
}

