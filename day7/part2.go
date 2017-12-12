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

func (p *program) String() string {
	return "prog{" + p.name + "}"
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
	byName := make(map[string]*program)
	parents := make(map[string]*program)
	for _, line := range input {
		if line == "" {
			continue
		}
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
		byName[prog.name] = &prog
	}
	for _, prog := range programs {
		prog.parent = parents[prog.name]
	}
	var root *program
	for _, prog := range programs {
		if prog.parent == nil {
			if root != nil {
				panic("shouldn't get here")
			}
			root = prog
		}
	}
	var treeSum func(*program) int
	found := false
	treeSum = func(p *program) int {
		if len(p.children) == 0 {
			return p.weight
		}
		sum := p.weight
		balanced := make(map[int][]*program)
		for _, name := range p.children {
			child := byName[name]
			csum := treeSum(child)
			sum += csum
			balanced[csum] = append(balanced[csum], child)
		}
		if len(balanced) != 1 && !found {
			found = true
			var target int
			var csum int
			var unbalanced *program
			for weight, children := range balanced {
				if len(children) == 1 {
					unbalanced = children[0]
					csum = weight
				} else {
					target = weight
				}
			}
			fmt.Println(unbalanced.name, unbalanced.weight, target, csum, unbalanced.weight-(csum-target))
		}
		return sum
	}
	treeSum(root)
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return int(n)
}
