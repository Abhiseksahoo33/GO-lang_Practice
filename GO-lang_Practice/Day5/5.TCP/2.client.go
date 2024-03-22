package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Tcp client")

	con, err := net.Dial("tcp", ":1235")
	if err != nil {
		fmt.Println("Received Connectrion error", err)
	} else {
		con.Write([]byte("Hello This is from our test client"))
	}

}
