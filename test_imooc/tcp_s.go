package main

import (
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		return
	}

	defer listen.Close()

	for {
		conn, err := listen.Accept()
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

		log.Print("server recv : ", string(buf[0:len]))

		conn.Write([]byte("server"))
	}
}
