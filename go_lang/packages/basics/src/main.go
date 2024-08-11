package main

import (
	"errors"
	"fmt"
	"strconv"
)

// naked return, we return a variable by its name.
func greetWithNameAndAge(name string, age int) (greeting string) {
	greeting = "Hello, my name is " + name + " and I'm " + strconv.Itoa(age) + " years old"
	return
}

// introducing errors
func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by 0")
	}
	return a / b, nil
}

func main() {
	fmt.Println(greetWithNameAndAge("Marouane", 23))
	fmt.Println(divide(10, 2))
	fmt.Println(divide(20, 0))
}
