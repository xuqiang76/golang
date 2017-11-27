package main

import (
    "log"
    "net"
    "time"
)

func main() {
    conn, err := net.Dial("tcp", "127.0.0.1:1234")
    if err != nil {
        log.Print("client dial : ", err)
        return
    }

    go send(conn)

    buf := make([]byte, 1024)
    for {
        len, err := conn.Read(buf)
        if err != nil {
            conn.Close()
            log.Print("client read : ", err)
            return
        }
        log.Print(string(buf[0:len]))
    }
}

func send(conn net.Conn) {
    for {
        time.Sleep(3 * time.Second)
        _, err := conn.Write([]byte("client send"))
        if err != nil {
            log.Print("client send : ", err)
            conn.Close()
            break
        }
    }
}