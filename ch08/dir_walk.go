/*
    The function you pass to Walk is called for every file and folder in the
    root folder. The function takes three arguments:
    1) path, which is the path to the file
    2) info, which is the information for the file (the same information
       you get from using os.Stat), and
    3) err, which is any error that was received while walking the directory.
    The function returns an error and you can return filepath.SKipDir to stop
    walking immediately.
*/
package main

import (
    "fmt"
    "os"
    "path/filepath"
)

func main() {
    d := ".."
    filepath.Walk(d, func(path string, info os.FileInfo, err error) error {
        fmt.Println(path)
        return nil
    })
}

