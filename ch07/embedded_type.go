package main
import "fmt"

type Person struct {
    Name string
}
type Android struct {
    Person
    Model string
}

func (p *Person) Talk() {
    fmt.Println("Hi, my name is ", p.Name)
}

func main() {
    // The type "Person" in "Android" does not have a name.
    // So "Person" is an embedded type.
    // The Person struct can be accessed using the type name.
    a := new(Android)
    a.Person.Name = "Inigo Montoya"
    a.Person.Talk()
    // We can also call any Person method directly on the Android.
    a.Talk()
}

