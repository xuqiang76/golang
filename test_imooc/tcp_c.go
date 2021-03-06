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

	defer conn.Close()
	go c_send(conn)

	buf := make([]byte, 1024)
	for {
		len, err := conn.Read(buf)
		if err != nil {
			return
		}

		log.Println("client recv : ", string(buf[0:len]))
	}
}

func c_send(conn net.Conn) {
	for {
		time.Sleep(3 * time.Second)
		_, err := conn.Write([]byte("client"))
		if err != nil {
			return
		}
	}
}
