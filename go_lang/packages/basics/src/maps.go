package main

import "fmt"

type person struct {
	first, last string
}

type author struct {
	penName string
	person
}

func main() {

	// declare a map: key_type = string, value_type = author
	var authors map[string]author
	// init the map
	authors = make(map[string]author)
	authors["MB"] = author{
		person: person{
			first: "Marouane",
			last:  "Boufarouj",
		},
		penName: "Bic",
	}
	authors["YB"] = author{
		person: person{
			first: "Younes",
			last:  "Boufarouj",
		},
		penName: "Bic",
	}

	fmt.Printf("%#v\n", authors)

	// read a value with a known key
	fmt.Println("MB : ", authors["MB"].person.first, authors["MB"].person.last)

	// delete a key from a map
	delete(authors, "YB")
	fmt.Printf("%#v\n", authors)

	_, ok = authors["YB"]

	if ok {
		fmt.Println("Key exists")
	}
}
