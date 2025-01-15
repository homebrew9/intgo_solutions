package main

import (
    "fmt"
    "os"
)

func main() {
    // To get the contents of a directory, we use the os.Open function but
    // give it a directory path instead of a file name. Then we call the
    // Readdir method.
    dir, err := os.Open(".")
    if err != nil {
        return
    }
    defer dir.Close()

    // Readdir takes a single argument that limits the number of entries
    // returned. By passing in -1, we return all of the entries.
    fileInfos, err := dir.Readdir(-1)
    if err != nil {
        return
    }
    for _, fi := range fileInfos {
        fmt.Println(fi.Name())
    }
}

