package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Read file..")
	f, e := os.Open("C:/Users/sahooab/Downloads/AMF-5792_SBR_Show_n_Exec/FRN-16301 Source Based Routing v0.1 - Copy.txt")
	if e != nil {
		fmt.Println(e)
	}
	defer f.Close()
	data := make([]byte, 10, 10) //Read only 10 .. 10 bytes
	// str := strings.Builder{}
	for {
		n, e := f.Read(data)
		if e != nil {
			fmt.Println(e)
		} else {
			fmt.Println("Read", n, "Bytes")
			fmt.Println(string(data))
			// str += string(data)
			if n != 10 {
				break

			}
		}
	}
	// fmt.Println(str.string())
}

/*
tes of Uti
ls
        o For
mat the JS
ON - ** Ne
ed sample
format fro
m SGSN tea
m
        o Send
 the respo
*/
