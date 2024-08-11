package main

import (
	"fmt"
)

type author struct {
	first, last string
}

// this is a method related to instance of author (a is a receiver in this case)
func (a author) getFullName() string {
	return a.first + " " + a.last
}

// this is a normal standalone function
func get_Full_Name(a author) string {
	return a.first + " " + a.last
}

func main() {
	a := author{
		first: "Jhonny",
		last:  "Deep",
	}

	fmt.Printf("My name is %s\n", a.getFullName())
	// just optional (for the sake of understanding!)
	fmt.Printf("My name is %s\n", get_Full_Name(a))

}
