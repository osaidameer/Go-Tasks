package main

import "fmt"

type employee struct {

	name string
	salary int	
	position string

}
	
type company struct {	
	companyName string
	employees []employee

}

func main(){
	emp1 := employee{"Osaid", 0, "Student"}
	emp2 := employee{"Talha", 0, "Student"}
	emp3 := employee{"Sahil", 0, "Student"}

	empArray := []employee{emp1, emp2, emp3}

	c1 := company{"Tetra", empArray}

	fmt.Println(c1)
}