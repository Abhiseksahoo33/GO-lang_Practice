package main

import "fmt"

type Std struct {
	Roll_no int
	Name    string
}

func InitStudent() *Std {
	fmt.Println("Inside InitStudent")
	return new(Std)
}

var (
	my_var *Std = InitStudent()
)

func init() {
	fmt.Println("This is init")
	my_var = InitStudent() //first global data segment will execute
}
func main() {
	fmt.Println("Inside ")
}

/*
Inside InitStudent
This is init
Inside InitStudent
Inside
*/
