/*
   This example is very similar to the crc32 one, because both crc32 and sha1
   implement the hash.Hash interface. The main difference is that whereas crc32
   computes a 32-bit hash, sha1 computes a 160-bit hash. There is no native type
   to represent a 160-bit number, so we use a slice of 20 bytes instead.
*/
package main
import (
    "fmt"
    "crypto/sha1"
    "strings"
)

func main() {
    h := sha1.New()
    h.Write([]byte("test"))
    bs := h.Sum([]byte{})
    fmt.Println("test")
    fmt.Println(bs)
    fmt.Println(strings.Repeat("=",110))
    fmt.Println("SHA1 hashes on all elements of an array...")
    fmt.Println(strings.Repeat("=",110))
    slice := []string{
        "The quick brown fox jumps over the lazy dog",
        "the quick brown fox jumps over the lazy dog",
        "The woods are lovely, dark and deep",
        "Now is the time for all young men",
        "abracadabra",
    }
    for _, value := range(slice) {
        fmt.Println("String      = ", value)
        h.Write([]byte(value))
        bs = h.Sum([]byte{})
        fmt.Println("SHA1 Hash   = ", bs)
        fmt.Println(strings.Repeat("-",95))
    }
}

