package main

import "fmt"

type Student struct {
	RollNo int
	Name   string
	Class  string
	Score  Marks // embedding marks to students structure
}

type Marks struct {
	m1, m2, m3, m4, m5, m6 int
	total                  int
	avg                    float32
}

func (std Student) SetRollNo(roll_no int) Student {
	std.RollNo = roll_no
	return std
}

func (std Student) SetName(name string) Student {
	std.Name = name
	return std
}

func (std Student) SetClass(class string) Student {
	std.Class = class
	return std
}

func (std Student) SetM1(m1 int) Student {
	std.Score.m1 = m1
	return std
}

func (std Student) SetM2(m2 int) Student {
	std.Score.m2 = m2
	return std
}

func (std Student) SetM3(m3 int) Student {
	std.Score.m3 = m3
	return std
}

func main() {
	fmt.Println("Demo: Methods")

	s := Student{}

	s = s.SetRollNo(1).SetName("Abhisek Sahoo").SetClass("MCA").SetM1(100).SetM2(97).SetM3(88)
	fmt.Println(s)- //{1 Abhisek Sahoo MCA {100 97 88 0 0 0 0 0}}
}
