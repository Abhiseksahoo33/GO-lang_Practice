// server.go
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
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		// Receive data from the client
		buf := make([]byte, 1024)
		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println("Read ", n, "bytes", "Contents:\n", string(buf))
		}

		// Send data back to the client
		_, err = conn.WriteToUDP(buf[:n], addr)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	// Close the connection
	conn.Close()
}
