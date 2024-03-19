package main

import "fmt"

type Comp struct {
	ModelNo string
	Name    string
}

func main() {
	fmt.Println("Pointers...")

	v := Comp{ModelNo: "12133", Name: "Hp"}
	fmt.Println(v)

	fmt.Println(v.ModelNo, v.Name) //12133 Hp

	ptrComp := &v

	fmt.Println(ptrComp)
	fmt.Println(ptrComp.ModelNo, ptrComp.Name) //12133 Hp

	var ptrC *Comp = &v
	fmt.Println(ptrC) //&{12133 Hp}
}
