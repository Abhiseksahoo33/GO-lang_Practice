package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Product catlog server..")

	//Initialize get handler
	err := http.ListenAndServe(":8080", CraateCatlogMux())
	if err != nil {
		log.Fatalln("Error while creating server", err)
	}

}
