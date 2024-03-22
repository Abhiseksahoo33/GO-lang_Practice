// client.go
package main

import (
	"fmt"
	"net"
)

func main() {
	// Create a UDP address
	addr, err := net.ResolveUDPAddr("udp", ":1234")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create a UDP connection
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Send data to the server
	_, err = conn.Write([]byte("Hello, world!"))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Receive data from the server
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the data from the server
	fmt.Println(string(buf[:n]))

	// Close the connection
	conn.Close()
}
