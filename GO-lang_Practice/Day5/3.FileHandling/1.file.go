package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
	fmt.Println("File handling.. Fs.Walk")
	root := "C:/Users/sahooab/Downloads/AMF-5792_SBR_Show_n_Exec"
	fsm := os.DirFS(root)
	fs.WalkDir(fsm, ".", func(path string, d fs.DirEntry, err error) error {
		fmt.Println(path)
		info, err := d.Info()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(info)
		}
		return nil
	})
}
