package main
import (
    "encoding/gob"
    "fmt"
    "net"
)

// In Go, we can create a TCP server using the net package's Listen function.
// Listen takes a network type (in our case tcp) and an address and port to
// bind, and returns a net.Listener.
/*
type Listener interface {
    // Accept waits for and returns the next connection to the listener
    Accept() (c Conn, err error)
    // Close closes the listener.
    // Any blocked Accept operations will be unblocked and return errors.
    Close() error
    // Addr returns the listener's network address
    Addr() Addr
}
*/

// Once we have a Listener, we call Accept, which waits for a client to connect
// and returns a net.Conn. A net.Conn implements the io.Reader and io.Writer
// interfaces, so we read from it and write to it just like a file.
func server() {
    // listen to a port
    ln, err := net.Listen("tcp", ":9999")
    if err != nil {
        fmt.Println(err)
        return
    }
    for {
        // accept a connection
        c, err := ln.Accept()
        if err != nil {
            fmt.Println(err)
            continue
        }
        // handle the connection
        go handleServerConnection(c)
    }
}

// This example uses the encoding/gob package, which makes it easy to encode Go
// value so that other Go programs (or the same Go program, in this case) can
// read them. Additional encodings are availabe in packages underneath encoding
// (like encoding/json) as well as in third-party packages (e.g. we could use
// labix.org/v2/mgo/bson for bson support).
func handleServerConnection(c net.Conn) {
    // receive the message
    var msg string
    err := gob.NewDecoder(c).Decode(&msg)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("Received", msg)
    }
    c.Close()
}

func client() {
    // connect to the server
    c, err := net.Dial("tcp", "127.0.0.1:9999")
    if err != nil {
        fmt.Println(err)
        return
    }

    // send the message
    msg := "Hello, World"
    fmt.Println("Sending", msg)
    err = gob.NewEncoder(c).Encode(msg)
    if err != nil {
        fmt.Println(err)
    }
    c.Close()
}

func main() {
    go server()
    go client()

    var input string
    fmt.Scanln(&input)
}

