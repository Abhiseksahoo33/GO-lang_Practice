package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Reading file..")
	data, err := os.ReadFile("C:/Users/sahooab/Downloads/AMF-5792_SBR_Show_n_Exec/FRN-16301 Source Based Routing v0.1.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(data)) // Problem will happen if size is too big
	}

}
