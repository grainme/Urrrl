package main

import (
	"fmt"
)

func main() {

	// let's create an array
	var a [3]int
	a[0] = 3 // [3 0 0]
	fmt.Println(a)

	// now we create a slice
	names := make([]string, 0)
	names = append(names, "Marouane")
	names = append(names, "Hitman")
	names = append(names, "Goodman")

	fmt.Println(names)
	fmt.Println(names[1:3])
	fmt.Println(names[:2])

}
