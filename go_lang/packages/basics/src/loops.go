package main

import "fmt"

func main() {

	// normal for loop
	s := "Hello"
	for i := 0; i < len(s); i++ {
		fmt.Println("at " , i , " : ", string(s[i]))
	}

	// while loop
	var n int32 = 5
	for n > 0 {
		fmt.Println(n)
		n--
	}

	// for-range loop : you can iterate whatever structure you have!
	arr := []int{100, 200, 300}

	for i, ele := range arr{
		fmt.Println(i, ele)
	}

	mpp := map[int]string{1: "one", 2: "two", 3: "three"}
	
	for first, second := range mpp{
		fmt.Println(first, second)
	}
}
