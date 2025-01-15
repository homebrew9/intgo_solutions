package main
import "fmt"

func average(xs []float64) float64 {
    //panic("Not implemented!")
    total := 0.0
    for _, value := range xs {
        total += value
    }
    return total / float64(len(xs))
}

func f2() (r int) {
    r = 1
    return
}

func f3() (int, int) {
    return 23, 37
}

func add(args ...int) int {
    total := 0
    for _, v := range args {
        total += v
    }
    return total
}

func main() {
    arr := []float64{1.23, 2.34, 3.45, 4.56, 5.67, 6.78}
    //var retval float64
    //retval = average(arr)
    //fmt.Println(retval)
    fmt.Println("arr = ", arr)
    fmt.Println("avg = ", average(arr))
    fmt.Println("==================================================")
    fmt.Println("Return types can have names")
    fmt.Println(f2())
    fmt.Println("==================================================")
    fmt.Println("Multiple values can be returned")
    //var x, y int
    //x, y = f3()
    x, y := f3()
    fmt.Println(x, y)
    fmt.Println("==================================================")
    fmt.Println("Example of a variadic function")
    fmt.Println(add(1,2,3))
    fmt.Println("==================================================")
    fmt.Println("Passing a slice to a variadic function")
    xs := []int{20,30,40}
    fmt.Println(xs)
    fmt.Println(add(xs...))
    fmt.Println("==================================================")
    fmt.Println("Passing an array to a variadic function")
    a := [3]int{50,51,52}
    fmt.Println(a)
    // Errors out: "cannot use a (type [3]int) as type []int in argument to add"
    //fmt.Println(add(a...))
    fmt.Println("Does not work!")
}

