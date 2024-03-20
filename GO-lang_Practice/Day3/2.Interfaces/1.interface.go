package main

import "fmt"

/*
interface basically scalaton of methods, objects are compilant to that interface
type <Name of interface> interface{
	Method1()
	Method2()
}
*/

type Emp struct {
	Id   int
	Name string
}

func (e Emp) Print() {
	fmt.Println("eID: ", e.Id)
	fmt.Println("eName: ", e.Name)
}

func (e Emp) String() string { // Reflection of interface implimentation
	return fmt.Sprint("eID: ", e.Id, "eName: ", e.Name)
}

type Std struct {
	id   int
	Name string
}

func (s Std) Print() {
	fmt.Println("sID: ", s.id)
	fmt.Println("sName: ", s.Name)
}

func (s Std) String() string { // Reflection of interface implimentation
	return fmt.Sprint("sID: ", s.id, "sName: ", s.Name)
}

type Printer interface {
	Print()
}

func Print(p Printer) {
	p.Print()
}

func main() {
	fmt.Println("Exploring interfaces..")
	var e, s Printer

	e = Emp{Id: 1, Name: "Abhisek"}
	s = Std{id: 2, Name: "Itu"}

	e.Print()
	s.Print()

	// print(e)
	// print(s)

	/*eID:  1
	  eName:  Abhisek
	  sID:  2
	  sName:  Itu*/

	fmt.Println(e) // Reflection of interface
	fmt.Println(s)
	/*eID: 1eName: Abhisek
	sID: 2sName: Itu*/

}
