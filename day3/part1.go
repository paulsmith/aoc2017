package main

import (
	"fmt"
)

func main() {
	var input int
	if _, err := fmt.Scanf("%d", &input); err != nil {
		panic(err)
	}
	i := 1
	var square, prev int
	for {
		square = i * i
		if prev < input && input < square {
			break
		}
		i += 2
		prev = square
	}
	fmt.Println(input, i, prev, square)
	diff := square - input
	fmt.Println(((i / 2) - diff) + (i / 2))
}
