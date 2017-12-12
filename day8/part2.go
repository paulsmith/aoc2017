package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type inst struct {
	reg    string
	dir    string
	amount int
	cond   cond
}

type cond struct {
	reg string
	op  string
	n   int
}

func main() {
	b, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	var instructions []inst
	registers := make(map[string]int)
	lines := strings.Split(string(b), "\n")
	for _, line := range lines {
		bits := strings.Split(line, " ")
		inst := inst{
			reg:    bits[0],
			dir:    bits[1],
			amount: atoi(bits[2]),
			cond: cond{
				reg: bits[4],
				op:  bits[5],
				n:   atoi(bits[6]),
			},
		}
		instructions = append(instructions, inst)
		registers[inst.reg] = 0
		registers[inst.cond.reg] = 0
	}
	var max int
	for pc := 0; pc < len(instructions); pc++ {
		inst := instructions[pc]
		if inst.dir == "dec" {
			inst.amount *= -1
		}
		cond := inst.cond
		reg := cond.reg
		doit := false
		switch cond.op {
		case ">":
			doit = registers[reg] > cond.n
		case "<":
			doit = registers[reg] < cond.n
		case ">=":
			doit = registers[reg] >= cond.n
		case "<=":
			doit = registers[reg] <= cond.n
		case "==":
			doit = registers[reg] == cond.n
		case "!=":
			doit = registers[reg] != cond.n
		default:
			panic(cond.op)
		}
		if doit {
			registers[inst.reg] += inst.amount
		}
		if registers[inst.reg] > max {
			max = registers[inst.reg]
		}
	}
	fmt.Println(max)
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return int(n)
}
