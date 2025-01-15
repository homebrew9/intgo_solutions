package main
import "fmt"

func get_line(s string, n int) string {
    var line string = ""
    for i := 1; i <= n; i++ {
        line += s
    }
    return line
}

func main() {
    //var x map[string]int
    x := make(map[string]int)
    x["key"] = 10
    fmt.Println(x)
    fmt.Println(x["key"])
    fmt.Println("==================================================")
    fmt.Println("Trying out for..range construct for maps")
    x["apples"] = 12
    x["bats"] = 23
    x["cars"] = 5
    x["doors"] = 2
    for k, v := range x {
        fmt.Printf("%-10s => %3d\n", k, v)
    }
    fmt.Println(">>map len =", len(x))
    fmt.Println("==================================================")
    elements := make(map[string]string)
    elements["H"] = "Hydrogen"
    elements["He"] = "Helium"
    elements["Li"] = "Lithium"
    elements["Be"] = "Beryllium"
    elements["B"] = "Boron"
    elements["C"] = "Carbon"
    elements["N"] = "Nitrogen"
    elements["O"] = "Oxygen"
    elements["F"] = "Fluorine"
    elements["Ne"] = "Neon"
    fmt.Println("elements = ", elements)
    fmt.Println("Li = ", elements["Li"])
    fmt.Println("Un = ", elements["Un"])
    name, ok := elements["Un"]
    fmt.Println(name, ok)
    // The following is an interesting construct!
    if name, ok := elements["Un"]; ok {
        fmt.Println(name, ok)
    } else {
        fmt.Println("Key not found: Un")
    }
    fmt.Println(">>> Looping through a list of potential keys")
    // Try a list of keys
    for _, value := range([4]string{"Un","C","No","Ne"}) {
        //fmt.Println(value)
        if name, ok := elements[value]; ok {
            fmt.Println("Key found    : ", name, ok)
        } else {
            fmt.Println("Key not found: ", value)
        }
    }
    fmt.Println("==================================================")
    // Another way to declare maps
    capitals := map[string]string {
        "USA" : "Washington D.C.",
        "Germany" : "Berlin",
        "Japan" : "Tokyo",
    }
    fmt.Println(capitals)
    fmt.Println("==================================================")
    // More complex map
    elem := map[string]map[string]string{
        "H" : map[string]string{
            "name" : "Hydrogen",
            "state" : "gas",
        },
        "He" : map[string]string{
            "name" : "Helium",
            "state" : "gas",
        },
        "Li" : map[string]string{
            "name" : "Lithium",
            "state" : "solid",
        },
        "Be" : map[string]string{
            "name" : "Beryllium",
            "state" : "solid",
        },
        "B" : map[string]string{
            "name" : "Boron",
            "state" : "solid",
        },
        "C" : map[string]string{
            "name" : "Carbon",
            "state" : "solid",
        },
        "N" : map[string]string{
            "name" : "Nitrogen",
            "state" : "gas",
        },
        "O" : map[string]string{
            "name" : "Oxygen",
            "state" : "gas",
        },
        "F" : map[string]string{
            "name" : "Fluorine",
            "state" : "gas",
        },
        "Ne" : map[string]string{
            "name" : "Neon",
            "state" : "gas",
        },
    }
    var ch string = "-"
    fmt.Printf("%-6s|%-20s|%s\n", "Symbol", "Name", "State at room temp")
    fmt.Printf("%-6s+%-20s+%s\n", get_line(ch,6), get_line(ch,20), get_line(ch,30))
    for k, v := range(elem) {
        fmt.Printf("%-6s|%-20s|%s\n", k, v["name"], v["state"])
    }
    fmt.Printf("%-6s+%-20s+%s\n", get_line(ch,6), get_line(ch,20), get_line(ch,30))
}

