package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var mem []int
	for scanner.Scan() {
		jump := atoi(scanner.Text())
		mem = append(mem, jump)
	}
	var pc, n int
	for pc < len(mem) {
		jump := mem[pc]
		mem[pc] += 1
		pc += jump
		n++
	}
	fmt.Println(n)
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return int(n)
}
