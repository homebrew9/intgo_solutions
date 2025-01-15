package main
import "fmt"

func main() {
    // "var x" creates an empty zero-length slice
    // "x := make()" creates a slice of length 5, of capacity 10, populated with 5 zeros
    //var x []float64
    x := make([]float64, 5, 10)
    fmt.Println(x)
    // Another way
    arr := [5]float64{1,2,3,4,5}
    y := arr[1:4]
    z := arr[:]
    fmt.Println("array arr = ", arr)
    fmt.Println("slice y   = ", y)
    fmt.Println("slice z   = ", z)

    fmt.Println("==================================================")
    fmt.Println("Testing append function")
    //slice1 := []int{1,2,3}
    slice1 := make([]int, 3, 5)
    slice1[0] = 1
    slice1[1] = 2
    slice1[2] = 3
    slice2 := append(slice1, 29, -71, 0.0, 3599, -191)
    fmt.Println("slice1 = ", slice1)
    fmt.Println("slice2 = ", slice2)

    fmt.Println("==================================================")
    fmt.Println("Testing copy function")
    slice3 := []int{1,2,3}
    slice4 := make([]int, 2)
    // copy(destination, source)
    // copy contents of slice3(source) into slice4(destination) - as much as possible.
    copy(slice4, slice3)
    fmt.Println("slice3 = ", slice3)
    fmt.Println("slice4 = ", slice4)
}

