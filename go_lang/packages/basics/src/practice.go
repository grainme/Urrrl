package main

import (
	"fmt"
)

// we want to do a*b without *
func mult_(a int, b int) int {
	if a == 0 {
		return 0
	} else if a == 1 {
		return b
	} else if a%2 == 0 {
		res := mult_(a/2, b)
		return res + res
	}
  // 5 / 2 = 2 * B + 2 * B = 4*B + B
	res := mult_(a/2, b)
	return res + res + bx
}
ww
// method 2 : << shifting to the left (padding 0) == *2
// >> shifting to the right (padding 0) == /2
// we want to do a*b without *
func mult_2(a int, b int) (res int) {
	// a & 1 : 101 (5) & 001 = 001 = 1 (check if odd or even)
  for a > 0 {
    // if a is odd 
    if a&1 == 1 {
      res += b
    }
    // a == 0 is the stopping condition
    a = a >> 1
    // a * b = a/2 * b*2 - this is the trick :)
    b = b << 1
  }
  return
}


func main() {
	fmt.Println("5 * 2 = ", mult_2(5, 2))
  fmt.Println("0 * 2 = ", mult_2(0, 2))
  fmt.Println("1 * 7 = ", mult_2(1, 7))
  fmt.Println("4 * 9 = ", mult_2(4, 9))
}
