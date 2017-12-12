package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type program struct {
	name     string
	weight   int
	children []string
	parent   *program
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(b), "\n")
	nameWeightRe := regexp.MustCompile(`([a-z]+) \((\d+)\)`)
	aboveRe := regexp.MustCompile(` -> `)
	var programs []*program
	parents := make(map[string]*program)
	for _, line := range input {
		m := nameWeightRe.FindStringSubmatch(line)
		if len(m) != 3 {
			panic(fmt.Sprintf("expected 2 submatches, got %d", len(m)-1))
		}
		prog := program{
			name:   m[1],
			weight: atoi(m[2]),
		}
		loc := aboveRe.FindStringIndex(line)
		if loc != nil {
			prog.children = strings.Split(line[loc[0]+4:], ", ")
			for _, child := range prog.children {
				parents[child] = &prog
			}
		}
		programs = append(programs, &prog)
	}
	for _, prog := range programs {
		prog.parent = parents[prog.name]
	}
	for _, prog := range programs {
		if prog.parent == nil {
			fmt.Println(prog.name)
		}
	}
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return int(n)
}
