package main

import (
	"fmt"
)

func parseOddsEvens(ints []int) (evens []int, odds []int) {
	for _, ele := range ints {
		if ele%2 == 0 {
			evens = append(evens, ele)
		} else {
			odds = append(odds, ele)
		}
	}
	return
}

func main() {
	ints := []int{1, 2, 3, 4, 5, 6, 7, 8}
	evens, odds := parseOddsEvens(ints)

	fmt.Println(evens)
	fmt.Println(odds)
}
