package main

import "fmt"

var val string = "value of val"

func update_global_val() {
	val = "Marouane is cooking"
}

func main() {
	val := 10
	fmt.Printf("this is the type of val (%T) and its value %d \n", val, val)
}
