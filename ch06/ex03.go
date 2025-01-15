package main
import (
    "fmt"
    "math"
    "math/rand"
    "time"
)

func greatest_flt(x ...float64) float64 {
    var greatest_value float64
    for i, value := range x {
        if (i == 0) || (value > greatest_value) {
            greatest_value = value
        }
    }
    return greatest_value
}

func greatest_int(x ...int) int {
    var greatest_value int
    for i, value := range x {
        if (i == 0) || (value > greatest_value) {
            greatest_value = value
        }
    }
    return greatest_value
}

func RoundPlus(f float64, places int) (float64) {
    // Copied from https://gist.github.com/DavidVaini/10308388
    shift := math.Pow(10, float64(places))
    return math.Round(f * shift) / shift;
}

func main() {
    fmt.Println(greatest_flt(1.1,2.2,3.3))
    fmt.Println("===================================================================================")
    fmt.Println("Testing with slice of random float64 numbers...")
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    slice_flt := make([]float64, 10)
    iterations := 5
    for iterations > 0 {
        for i := 0; i < 10; i++ {
            slice_flt[i] = RoundPlus(r.Float64(), 3)
        }
        fmt.Println("slice_flt    = ", slice_flt)
        fmt.Println("greatest_flt = ", greatest_flt(slice_flt...))
        fmt.Println("<<==========>>")
        iterations--
    }
    fmt.Println("===================================================================================")
    fmt.Println("Testing with slice of random ints...")
    slice_int := make([]int, 10)
    iterations = 5
    for iterations > 0 {
        for i := 0; i < 10; i++ {
            slice_int[i] = r.Intn(1000)
        }
        fmt.Println("slice_int    = ", slice_int)
        fmt.Println("greatest_int = ", greatest_int(slice_int...))
        fmt.Println("<<==========>>")
        iterations--
    }
}

