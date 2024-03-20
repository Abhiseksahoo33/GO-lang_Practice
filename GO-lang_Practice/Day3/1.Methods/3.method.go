package main

import "fmt"

type Compensation struct {
	Basic, Hra, Allowance float64
	Tax, Pf               float64
}

type Employee struct {
	EmpId int
	Name  string
	Dept  string
	Compensation
}

func (emp *Employee) Set(edid int, name, dept string) {
	emp.EmpId = edid
	emp.Name = name
	emp.Dept = dept
}

func (cmp *Compensation) SetSalary(basic, hra, allowance float64, tax, pf float64) {
	cmp.Basic = basic
	cmp.Hra = hra
	cmp.Allowance = allowance
	cmp.Tax = tax
	cmp.Pf = pf
}

func (cmp *Compensation) Print() {
	total := (cmp.Basic + cmp.Allowance + cmp.Hra) - (cmp.Tax + cmp.Pf)
	fmt.Println(cmp)
	fmt.Print("Net take home: ", total)
}

func (emp *Employee) Print() {
	fmt.Println(emp)
}

func main() {
	fmt.Println("Methods...")
	emp := new(Employee)
	emp.Set(1, "Abhisek", "AMF")
	emp.SetSalary(100000.0, 10000.0, 100000.0, 40000.0, 12000.0)
	emp.Print()
}
