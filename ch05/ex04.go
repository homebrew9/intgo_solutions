package main
import "fmt"

func main() {
    x := []int{
        48,96,86,68,
        57,82,63,70,
        37,34,83,27,
        19,97, 9,17,
    }
    // Code to find the smallest number in the list above
    var smallest int
    for i, value := range x {
        //if i == 0 {
        //    smallest = value
        //} else if value < smallest {
        //    smallest = value
        //}
        if (i == 0) || (value < smallest) {
            smallest = value
        }
    }
    fmt.Println(smallest)
}

