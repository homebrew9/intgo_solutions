package main

import ("fmt"; "sort")

type Person struct {
    Name string
    Age int
}
type ByName []Person

func (ps ByName) Len() int {
    return len(ps)
}
func (ps ByName) Less(i, j int) bool {
    return ps[i].Name < ps[j].Name
}
func (ps ByName) Swap(i, j int) {
    ps[i], ps[j] = ps[j], ps[i]
}

type ByAge []Person
func (this ByAge) Len() int {
    return len(this)
}
func (this ByAge) Less(i, j int) bool {
    return this[i].Age < this[j].Age
}
func (this ByAge) Swap(i, j int) {
    this[i], this[j] = this[j], this[i]
}

func main() {
    /*
      The Sort function in sort takes a sort.Interface and sorts it. The
      sort.Interface requires three methods: Len, Less and Swap.
      Len should return the length of the thing we are sorting.
      Less is used to determine whether item[i] is less than item[j]
      Swap swaps the items.
      To define our own sort, we create a new type (ByName) and make it
      equivalent to a slice of what we want to sort. We then define the three
      methods.
    */
    kids := []Person{
        {"Dave", 8},
        {"Alex", 11},
        {"Emma", 7},
        {"Cody", 9},
        {"Bill", 10},
    }
    fmt.Println("Kids           : ", kids)
    sort.Sort(ByName(kids))
    fmt.Println("Sorted by name : ", kids)
    sort.Sort(ByAge(kids))
    fmt.Println("Sorted by age  : ", kids)
}

