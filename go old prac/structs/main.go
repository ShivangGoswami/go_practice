package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	//alex := person{firstName: "Alex", lastName: "Anderson"}
	//fmt.Println(alex)
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex) //%+v for variable too %v only prints values
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	// jimPointer := &jim
	// jimPointer.updateName("jimmy")
	// jim.print()
	jim.updateName("jimmy")
	jim.print()
	//qwerty()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstname string) {
	(*p).firstName = newFirstname
}
