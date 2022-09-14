package main

import (
	"fmt"
	"strings"
)

type Student struct {
	rollnumber int 
	name string 
	address string
}
	
func NewStudent (rollno int, name string, address string) *Student {
	s := new(Student) 
	s.rollnumber = rollno
	s.name = name
	s.address = address
	return s
}

type StudentList struct {
	list []*Student
}

func (ls *StudentList) CreateStudent (rollno int, name string, address string) *Student {
	st := NewStudent (rollno, name, address) 
	ls.list = append(ls.list, st) 
	return st
}

func main() { 
	student := new(StudentList)
	student.CreateStudent (24, "Asim", "AAAAAA") 
	student.CreateStudent (25, "Naveed", "BBBBBB")
	student.CreateStudent(26, "Osaid", "CCCCCC")

	for i:=0; i<len(student.list); i++ {
		fmt.Printf("\n%s List %d %s", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		fmt.Printf("\nstudent rollno: %d \nstudent name: %s \nstudent address: %s", student.list[i].rollnumber, student.list[i].name, student.list[i].address)
	}
}	
	