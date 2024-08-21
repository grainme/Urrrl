package main

import "fmt"

func main() {

	fmt.Println("working from main...")

	// recovering (this is anonymous function btw)
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from panic: ", err)
		}
	}()

	// panic
	panic("something bad happened")

	fmt.Println("working from main...")
}

/*
	this is kind of like a try-catch mechanism but not precisely :)
*/
