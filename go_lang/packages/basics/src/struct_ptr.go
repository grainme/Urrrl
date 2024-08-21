// package + import of modules
package main

import (
	"fmt"
)

// structs
type person struct {
	first, last string
}
type author struct {
	penName string
	person
}

// this is a method related to instance of author (a is a receiver in this case)
func (a person) getFullName() string {
	return a.first + " " + a.last
}

func (a author) getFullName() string {
	return fmt.Sprintf("%s (%s)", a.person.getFullName(), a.penName)
}

// Main function
func main() {
	a := author{
		person: person{
			first: "Marouane",
			last:  "47",
		},
		penName: "Bic",
	}
	fmt.Println(a.getFullName())
}
