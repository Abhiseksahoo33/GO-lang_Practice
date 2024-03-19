package main

import (
	"fmt"
	"os"
)

/*
open a file
check if error occures, if yes return error
show the contents of the file
*/

func main() {
	filename := "go.txt"

	// Open the file for reading
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	//Read the entire contents of the file
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File contents:", string(data))
}
