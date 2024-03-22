package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Loggers..")

	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.SetPrefix("[Go Training]")
	log.Println("Entering main...")

	//log.Fatal("Exiting main") // will print hii then force exit

	log.Fatalln("Attempting fatalln")
	fmt.Println("Hello..") // will not execute

	log.Panic("panic msg")

}
