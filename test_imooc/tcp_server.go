package main

import (
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		return
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handle_conn(conn)
	}
}

func handle_conn(conn net.Conn) {
	buf := make([]byte, 1024)
	for {
		len, err := conn.Read(buf)
		if err != nil {
			conn.Close()
			return
		}

		log.Print("server read : ", string(buf[0:len]))
		conn.Write([]byte("server"))
	}
}
