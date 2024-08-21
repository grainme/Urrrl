package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	t, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	input := strings.TrimSpace(t)

	number, err_conv := strconv.Atoi(input)
	if err_conv != nil {
		panic(err_conv)
	}

	for i := 0; i < number; i++ {
		var d int
		sm := 0
		fmt.Scanf("%d", &d)
		for d > 0 {
			sm += d % 10
			d /= 10
		}
		fmt.Println(sm)
	}
}
