/*
   The net/rpc (remote procedure call) and net/rpc/jsonrpc packages provide an
   easy way to expose methods so they can be invoked over a network (rather than
   just in the program running them).
   This program is similar to the TCP example, except now we created an object
   to hold all the methods we want to expose and we call the Negate method from
   the client.
*/
package main
import (
    "fmt"
    "net"
    "net/rpc"
)

type Server struct {}
func (this *Server) Negate(i int64, reply *int64) error {
    *reply = -i
    return nil
}

func server() {
    rpc.Register(new(Server))
    ln, err := net.Listen("tcp", ":9999")
    if err != nil {
        fmt.Println(err)
        return
    }
    for {
        c, err := ln.Accept()
        if err != nil {
            continue
        }
        go rpc.ServeConn(c)
    }
}

func client() {
    c, err := rpc.Dial("tcp", "127.0.0.1:9999")
    if err != nil {
        fmt.Println(err)
        return
    }
    var result int64
    slice := []int64{999, 123, 0, 1729, 65536, -789, -1}
    for _, value := range slice {
        //err = c.Call("Server.Negate", int64(999), &result)
        err = c.Call("Server.Negate", value, &result)
        if err != nil {
            fmt.Println(err)
        } else {
            fmt.Printf("Server.Negate(%5d) = %7d\n",value,result)
        }
    }
}

func main() {
    go server()
    go client()

    var input string
    fmt.Scanln(&input)
}

