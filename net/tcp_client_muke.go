package main

import (
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		return
	}

	go c_send(conn)

	buf := make([]byte, 1024)
	for {
		len, err := conn.Read(buf)
		if err != nil {
			conn.Close()
			return
		}
		log.Print(string(buf[0:len]))
	}
}

func c_send(conn net.Conn) {
	for {
		time.Sleep(3 * time.Second)
		_, err := conn.Write([]byte("client send "))
		if err != nil {
			conn.Close()
			break
		}
	}
}
