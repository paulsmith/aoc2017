package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(b), "\n")
	groups := make(map[int][]int)
	for _, line := range lines {
		bits := strings.Split(line, " <-> ")
		id := atoi(bits[0])
		for _, s := range strings.Split(bits[1], ", ") {
			groups[id] = append(groups[id], atoi(s))
		}
	}
	seen := make(map[int]bool)
	var walk func(int)
	walk = func(gid int) {
		seen[gid] = true
		for _, id := range groups[gid] {
			if !seen[id] {
				walk(id)
			}
		}
	}
	var total int
	for id := range groups {
		if !seen[id] {
			walk(id)
			total++
		}
	}
	fmt.Println(total)
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return int(n)
}
