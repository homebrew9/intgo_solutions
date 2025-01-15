package main
import (
    "fmt"
    "hash/crc32"
    "io"
    "os"
)

func getHash(filename string) (uint32, error) {
    // open the file
    f, err := os.Open(filename)
    if err != nil {
        return 0, err
    }
    // remember to always close opened files
    defer f.Close()

    // create a hasher
    h := crc32.NewIEEE()
    // copy the file into the hasher
    // - copy takes (dst, src) and returns (bytesWritten, error)
    _, err = io.Copy(h, f)
    // we don't care about how many bytes were written, but we do want to
    // handle the error
    if err != nil {
        return 0, err
    }
    return h.Sum32(), nil
}

func main() {
    h1, err := getHash("test1.txt")
    if err != nil {
        fmt.Println("ERROR while processing test1.txt, see below:")
        fmt.Println(err)
        return
    }
    h2, err := getHash("test2.txt")
    if err != nil {
        fmt.Println("ERROR while processing test2.txt, see below:")
        fmt.Println(err)
        return
    }
    fmt.Println(h1, h2, h1 == h2)
}

