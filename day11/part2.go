package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	steps := strings.Split(string(b), ",")
	pos := []int{0, 0} // q, r
	orig := []int{0, 0}
	var max int
	for _, step := range steps {
		step = strings.TrimSpace(step)
		// axial coords
		switch step {
		case "n":
			pos[0], pos[1] = pos[0]+0, pos[1]-1
		case "ne":
			pos[0], pos[1] = pos[0]+1, pos[1]-1
		case "se":
			pos[0], pos[1] = pos[0]+1, pos[1]+0
		case "s":
			pos[0], pos[1] = pos[0]+0, pos[1]+1
		case "sw":
			pos[0], pos[1] = pos[0]-1, pos[1]+1
		case "nw":
			pos[0], pos[1] = pos[0]-1, pos[1]+0
		default:
			panic(step)
		}
		nsteps := dist(pos, orig)
		if nsteps > max {
			max = nsteps
		}
	}
	fmt.Println(max)
}

func dist(start, orig []int) int {
	return (abs(start[0]-orig[0]) + abs(start[0]+start[1]-orig[0]-orig[1]) + abs(start[1]-orig[1])) / 2
}

func abs(n int) int {
	return int(math.Abs(float64(n)))
}
