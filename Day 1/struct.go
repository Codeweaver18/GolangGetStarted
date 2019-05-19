package main

import (
	"fmt"
)

//Person:: structs =>classes
type Person struct {
	Fname string
	Lname string
	Title string
}

//instattiating struct
func main() {

	var student = Person{"Matsumu", "Monday", "Engr"}
	fmt.Println(student)
	var staff = Person{}
	staff.Fname = "Chloe"
	staff.Lname = "Obrian, "
	staff.Title = "Senior Software Engineer"
	fmt.Println(staff)
}
