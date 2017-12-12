package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "11	11	13	7	0	15	5	5	4	4	1	1	7	1	15	11"
	const nbank = 16
	var banks [nbank]int
	for i, v := range strings.Split(s, "\t") {
		banks[i] = atoi(v)
	}
	var cycle int
	seen := make(map[[nbank]int]int)
	for {
		cycle++
		if cycle == 100000 {
			panic("still going after 100000 cycles")
		}
		most, max := -1, -1
		for i := 0; i < nbank; i++ {
			if banks[i] > max {
				max = banks[i]
				most = i
			}
		}
		if most == -1 {
			panic("could not find most")
		}
		blocks := banks[most]
		banks[most] = 0
		for i := most; blocks > 0; blocks-- {
			i = (i + 1) % nbank
			banks[i]++
		}
		if sawcycle, ok := seen[banks]; ok {
			fmt.Println("saw", banks, "before on cycle", sawcycle, "cycles ago:", cycle-sawcycle)
			break
		}
		seen[banks] = cycle
	}
	fmt.Println(cycle)
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return int(n)
}
