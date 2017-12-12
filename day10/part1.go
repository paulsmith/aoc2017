package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "230,1,2,221,97,252,168,169,57,99,0,254,181,255,235,167"
	//input := "3,4,1,5"
	var lengths []int
	for _, s := range strings.Split(input, ",") {
		lengths = append(lengths, atoi(s))
	}
	const size = 256
	list := make([]int, size)
	for i := 0; i < size; i++ {
		list[i] = i
	}
	pos := 0
	skip := 0
	for _, length := range lengths {
		sublist := sel(list, pos, length)
		rev := reverse(sublist)
		update(list, rev, pos)
		pos += length + skip
		skip++
	}
	fmt.Println(list[0] * list[1])
}

func update(s []int, t []int, pos int) {
	start := pos
	for i := 0; i < len(t); i++ {
		pos = (start + i) % len(s)
		s[pos] = t[i]
	}
}

func sel(s []int, pos int, length int) (res []int) {
	start := pos
	for i := 0; i < length; i++ {
		pos = (start + i) % len(s)
		res = append(res, s[pos])
	}
	return
}

func reverse(s []int) (res []int) {
	for i := len(s) - 1; i >= 0; i-- {
		res = append(res, s[i])
	}
	return
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
