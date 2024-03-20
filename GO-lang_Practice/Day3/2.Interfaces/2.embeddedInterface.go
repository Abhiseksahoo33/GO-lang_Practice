package main

import "fmt"

type Object struct {
	Id   int
	Desc string
}

func (o *Object) Read() {
	fmt.Println(o)
}

func (o *Object) write(id int, desc string) {
	o.Id = id
	o.Desc = desc
}

type Reader interface {
	Read()
}

type Writer interface {
	write(id int, desc string)
}

type ObjectIO interface {
	Reader
	Writer
}

func main() {
	fmt.Println("Embeded interfaces")

	var oIO ObjectIO
	oIO = &Object{}
	oIO.write(1, "Test Object")
	oIO.Read()

	var r Reader
	var w Writer
	r = oIO
	w = oIO
	w.write(1, "Dummy")
	r.Read()
}

/*
Embedded interfaces
&{1 Test Object}
&{1 Dummy}
*/
