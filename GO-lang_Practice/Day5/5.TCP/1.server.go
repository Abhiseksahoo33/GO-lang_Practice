package main

import (
	"fmt"
	"log"
	"net"
)

func HandleConnection(conn net.Conn) {
	data := make([]byte, 1024, 1024)
	n, err := conn.Read(data)
	if err != nil {
		fmt.Println("Received Error while reading n", err)
	} else {
		fmt.Println("Read ", n, "bytes", "Contents:", string(data))
	}
	conn.Close()
}

func main() {
	fmt.Println("Exploring TCP Server")

	l, e := net.Listen("tcp", ":1235")

	if e != nil {
		log.Fatalln("failed to listen", e)
	}
	for {
		c, e := l.Accept()
		if e != nil {
			fmt.Println("Failed to accept the connection", e)
		} else {
			go HandleConnection(c)
		}
	}

}
