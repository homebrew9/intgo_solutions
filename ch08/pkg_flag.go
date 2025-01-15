/*
   The flag package allows us to parse arguments and flags sent to our program.
   This program generates a number between 0 and 6. We can change the max value
   by sending a flag (-max=100) to the program.
*/
package main
import ("fmt"; "flag"; "math/rand"; "time")

func main() {
    // Define flags
    maxp := flag.Int("max", 6, "the max value")
    // Parse
    flag.Parse()
    // Generate a number between 0 and max
    //fmt.Println(rand.Intn(*maxp))

    // We use a new seed as per documentation and then generate max int
    // Also see ch06/ex03.go
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    fmt.Println(r.Intn(*maxp))
}

