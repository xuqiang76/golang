package main 

import (
    "net"
    "log"
)

func func main() {
    listen, err := net.Listen("tcp", "127.0.0.1:1234")
    if err != nil {
        return
    }

    for {
        conn, err := listen.Accept()
        if err != nil {
            continue
        }

        go handle_conn(conn)
    }

    defer func () {
        listen.Close()
    }()
}

func handle_conn(conn net.Conn) {
    buf := make([]byte, 1024)

    defer conn.Close()

    for {
        len, err := conn.Read(buf)
        if err != nil {
            return
        }

        log.Print("recv msg ", string(buf[0:len]))
        conn.Write([]byte("server write "))
    }
}