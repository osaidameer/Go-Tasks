package main

import "fmt"

type Person struct {
	name string
	age int
	id int 
	job string
}

func print(P Person) {
	fmt.Println("Name: ", P.name)
	fmt.Println("ID: ", P.id)
	fmt.Println("Age: ", P.age)
	fmt.Println("Job: ", P.job)
}

func main() {
	var P1 Person
	P1.name = "Osaid"
	P1.id = 1990
	P1.age = 22
	P1.job = "salary"
	print(P1)
}