package main

import "fmt"

type Student struct {
	Roll int
	Name string
}

/*
 	func (<Receiver>) MethodName(<Arguments>) Return Value{
	receiver -> Object
 }
*/

func (std Student) Print() {
	fmt.Println("\t\t\t-----------------------------------")
	fmt.Println("\t\t\tStudents details are: ")
	fmt.Println("\t\t\t-----------------------------------")
	fmt.Println("\t\t\tRoll No:", std.Roll, "Name: ", std.Name)
	fmt.Println("\t\t\t-----------------------------------")
}

func main() {
	fmt.Println("Exploring Methods...") //method is associated with object #Function is generic

	var std = Student{Roll: 1, Name: "Abhisek"}
	fmt.Println(std) //{1 Abhisek}

	std.Print()
}
