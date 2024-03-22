package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Write file")

	fname := "C:/Users/sahooab/Downloads/AMF-5792_SBR_Show_n_Exec/text.txt"
	err := os.WriteFile(fname, []byte("This is test"), 0777)

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	data, err := os.ReadFile(fname)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println(string(data))
	}

}
