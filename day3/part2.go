package main

import (
	"fmt"
	"os"
)

type grid map[complex128]int

func (g grid) neighbors(pos complex128) int {
	return g[pos+(-1+0i)] +
		g[pos+(-1+1i)] +
		g[pos+(0+1i)] +
		g[pos+(1+1i)] +
		g[pos+(1+0i)] +
		g[pos+(1+-1i)] +
		g[pos+(0+-1i)] +
		g[pos+(-1+-1i)]
}

func (g grid) left(pos, dir complex128) bool {
	return g[pos+(dir*(0+1i))] > 0
}

func main() {
	var input int
	_, err := fmt.Sscanf(os.Args[1], "%d", &input)
	if err != nil {
		panic(err)
	}
	g := grid{
		0 + 0i: 1,
	}
	dir := 1 + 0i
	pos := 1 + 0i
	for ; g[pos-dir] < input; pos += dir {
		g[pos] = g.neighbors(pos)
		if !g.left(pos, dir) {
			dir *= 0 + 1i // rotate left
		}
	}
	fmt.Println(g[pos-dir])
}
