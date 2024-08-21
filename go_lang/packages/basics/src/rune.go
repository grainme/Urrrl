/*
	the problem we're trying to explore is ()
	-> why indexing strings in Rust (and even Go) is dangerous.
	-> because strings in modern languages (not all of them) are encoded in UTF-8
	   which is a variable-length encoding, meaning that different characters can
	   take up different number of bytes, unlike Ascii chars that take up 1 byte with no
	   exception whatsoever.
	-> DISCLAIMER : In Go, strings are also encoded in UTF-8, but the language allows indexing into a string.
	-> and this ladies and gentelmen can have bad consequences.
*/

package main

import (
	"fmt"
)

func main() {
	// 世界 is non ascii
	str := "Hello, 世界"

	// accessing bytes by index
	fmt.Printf("Type of str %T \n", str)
	fmt.Println("Byte at index 7", str[7])

	// converting bytes to characters :/ rune is alias of int32
	runes := []rune(str) // this is equivalent to []int32(str)
	fmt.Printf("Type of str %T \n", runes)
	fmt.Println("Byte at index 7", string(runes[7]))

	// loop through the string by character (rune) - instead of accessing using indexing
	fmt.Println("Characters in the string:")
	for i, r := range str {
		fmt.Printf("Index %d: %c\n", i, r)
	}
}
